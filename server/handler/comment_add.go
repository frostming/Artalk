package handler

import (
	"strings"

	"github.com/ArtalkJS/Artalk/internal/core"
	"github.com/ArtalkJS/Artalk/internal/entity"
	"github.com/ArtalkJS/Artalk/internal/i18n"
	"github.com/ArtalkJS/Artalk/internal/log"
	"github.com/ArtalkJS/Artalk/internal/utils"
	"github.com/ArtalkJS/Artalk/server/common"
	"github.com/gofiber/fiber/v2"
)

type ParamsAdd struct {
	Name    string `form:"name"`
	Email   string `form:"email"`
	Link    string `form:"link"`
	Content string `form:"content" validate:"required"`
	Rid     uint   `form:"rid"`
	UA      string `form:"ua"`

	PageKey   string `form:"page_key" validate:"required"`
	PageTitle string `form:"page_title"`

	SiteName string
	SiteID   uint
}

type ResponseAdd struct {
	Comment entity.CookedComment `json:"comment"`
}

// @Summary      Comment Add
// @Description  Create a new comment
// @Tags         Comment
// @Param        name           formData  string  false  "the comment name"
// @Param        email          formData  string  false  "the comment email"
// @Param        link           formData  string  false  "the comment link"
// @Param        content        formData  string  true   "the comment content"
// @Param        rid            formData  string  false  "the comment rid"
// @Param        ua             formData  string  false  "the comment ua"
// @Param        page_key       formData  string  true   "the comment page_key"
// @Param        page_title     formData  string  false  "the comment page_title"
// @Param        site_name      formData  string  false  "the site name of your content scope"
// @Security     ApiKeyAuth
// @Success      200  {object}  common.JSONResult{data=ResponseAdd}
// @Router       /add  [post]
func CommentAdd(app *core.App, router fiber.Router) {
	router.Post("/add", func(c *fiber.Ctx) error {
		var p ParamsAdd
		if isOK, resp := common.ParamsDecode(c, &p); !isOK {
			return resp
		}

		if strings.TrimSpace(p.Name) == "" {
			return common.RespError(c, i18n.T("{{name}} cannot be empty", Map{"name": i18n.T("Nickname")}))
		}
		if strings.TrimSpace(p.Email) == "" {
			return common.RespError(c, i18n.T("{{name}} cannot be empty", Map{"name": i18n.T("Email")}))
		}

		if !utils.ValidateEmail(p.Email) {
			return common.RespError(c, i18n.T("Invalid {{name}}", Map{"name": i18n.T("Email")}))
		}
		if p.Link != "" && !utils.ValidateURL(p.Link) {
			return common.RespError(c, i18n.T("Invalid {{name}}", Map{"name": i18n.T("Link")}))
		}

		var (
			ip      = c.IP()
			ua      = string(c.Request().Header.UserAgent())
			referer = string(c.Request().Header.Referer())
			isAdmin = common.CheckIsAdminReq(app, c)
		)

		// 允许传入修正后的 UA
		if p.UA != "" {
			ua = p.UA
		}

		// use site
		common.UseSite(c, &p.SiteName, &p.SiteID, nil)

		// find page
		page := app.Dao().FindCreatePage(p.PageKey, p.PageTitle, p.SiteName)

		// check if the user is allowed to comment
		if isAllowed, resp := common.CheckIsAllowed(app, c, p.Name, p.Email, page, p.SiteName); !isAllowed {
			return resp
		}

		// check reply comment
		var parentComment entity.Comment
		if p.Rid != 0 {
			parentComment = app.Dao().FindComment(p.Rid)
			if parentComment.IsEmpty() {
				return common.RespError(c, i18n.T("{{name}} not found", Map{"name": i18n.T("Parent comment")}))
			}
			if parentComment.PageKey != p.PageKey {
				return common.RespError(c, "Inconsistent with the page_key of the parent comment")
			}
			if !parentComment.IsAllowReply() {
				return common.RespError(c, i18n.T("Cannot reply to this comment"))
			}
		}

		// find user
		user := app.Dao().FindCreateUser(p.Name, p.Email, p.Link)
		if user.ID == 0 || page.Key == "" {
			log.Error("Cannot get user or page")
			return common.RespError(c, i18n.T("Comment failed"))
		}

		// update user
		user.Link = p.Link
		user.LastIP = ip
		user.LastUA = ua
		user.Name = p.Name // for 若用户修改用户名大小写
		user.Email = p.Email
		app.Dao().UpdateUser(&user)

		comment := entity.Comment{
			Content:  p.Content,
			PageKey:  page.Key,
			SiteName: p.SiteName,

			UserID: user.ID,
			IP:     ip,
			UA:     ua,

			Rid: p.Rid,

			IsPending:   false,
			IsCollapsed: false,
			IsPinned:    false,
		}

		// default comment type
		if !isAdmin && app.Conf().Moderator.PendingDefault {
			// 不是管理员评论 && 配置开启评论默认待审
			comment.IsPending = true
		}

		// save to database
		err := app.Dao().CreateComment(&comment)
		if err != nil {
			log.Error("Save Comment error: ", err)
			return common.RespError(c, i18n.T("Comment failed"))
		}

		// 异步执行
		go func() {
			// Page Update
			if app.Dao().CookPage(&page).URL != "" && page.Title == "" {
				app.Dao().FetchPageFromURL(&page)
			}

			// 垃圾检测
			if !isAdmin { // 忽略检查管理员
				// 同步执行
				if antiSpamService, err := core.AppService[*core.AntiSpamService](app); err == nil {
					antiSpamService.CheckAndBlock(&core.AntiSpamCheckPayload{
						Comment:      &comment,
						ReqReferer:   referer,
						ReqIP:        ip,
						ReqUserAgent: ua,
					})
				} else {
					log.Error("[AntiSpamService] err: ", err)
				}
			}

			// 通知发送
			if notifyService, err := core.AppService[*core.NotifyService](app); err == nil {
				if err := notifyService.Push(&comment, &parentComment); err != nil {
					log.Error("[NotifyService] notify push err: ", err)
				}
			} else {
				log.Error("[NotifyService] err: ", err)
			}
		}()

		cookedComment := app.Dao().CookComment(&comment)

		// IP 归属地
		if app.Conf().IPRegion.Enabled {
			if ipRegionService, err := core.AppService[*core.IPRegionService](app); err == nil {
				cookedComment.IPRegion = ipRegionService.Query(comment.IP)
			} else {
				log.Error("[IPRegionService] err: ", err)
			}
		}

		return common.RespData(c, ResponseAdd{
			Comment: cookedComment,
		})
	})
}
