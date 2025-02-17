package handler

import (
	"time"

	"github.com/ArtalkJS/Artalk/internal/core"
	"github.com/ArtalkJS/Artalk/server/common"
	"github.com/gofiber/fiber/v2"
)

// @Summary      User Logout
// @Description  Logout current user (applies to cookie identification only)
// @Description  User Logout
// @Tags         User
// @Success      200  {object}  common.JSONResult
// @Router       /logout  [post]
func UserLogout(app *core.App, router fiber.Router) {
	router.Post("/logout", func(c *fiber.Ctx) error {

		if !app.Conf().Cookie.Enabled {
			return common.RespError(c, "API cookie disabled")
		}

		if common.GetJwtStrByReqCookie(c) == "" {
			return common.RespError(c, "Not logged in yet, no need to log out")
		}

		// same as login, remove cookie
		setAuthCookie(app, c, "", time.Now().AddDate(0, 0, -1))

		return common.RespSuccess(c)
	})
}
