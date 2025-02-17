import type { ContextApi } from '~/types'
import Component from '../lib/component'
import * as Utils from '../lib/utils'
import CommentNode from '../comment/comment'
import ListHTML from './list.html?raw'
import ListLayout from './layout'
import { createComment as createCommentNode } from './comment'
import { initListPaginatorFunc } from './page'

export default class List extends Component {
  /** 列表评论集区域元素 */
  $commentsWrap!: HTMLElement
  public getCommentsWrapEl() { return this.$commentsWrap }

  protected commentNodes: CommentNode[] = []
  getCommentNodes() { return this.commentNodes }

  constructor (ctx: ContextApi) {
    super(ctx)

    // Init base element
    this.$el = Utils.createElement(ListHTML)
    this.$commentsWrap = this.$el.querySelector('.atk-list-comments-wrap')!

    // Init paginator
    initListPaginatorFunc(ctx)

    // Bind events
    this.initCrudEvents()
  }

  getListLayout() {
    return new ListLayout({
      $commentsWrap: this.$commentsWrap,
      nestSortBy: this.ctx.conf.nestSort,
      nestMax: this.ctx.conf.nestMax,
      flatMode: this.ctx.conf.flatMode as boolean,
      // flatMode must be boolean because it had been handled when Artalk.init
      createCommentNode: (d, c) => {
        const node = createCommentNode(this.ctx, d, c)
        this.commentNodes.push(node)
        return node
      },
      findCommentNode: (id) => this.commentNodes.find(c => c.getID() === id),
      getCommentDataList: () => this.ctx.getData().getComments(),
    })
  }

  private initCrudEvents() {
    this.ctx.on('list-load', (comments) => {
      // 导入数据
      this.getListLayout().import(comments)
    })

    this.ctx.on('list-loaded', (comments) => {
      if (comments.length === 0) this.$commentsWrap.innerHTML = ''
    })

    // When comment insert
    this.ctx.on('comment-inserted', (comment) => {
      this.getListLayout().insert(comment)
    })

    // When comment delete
    this.ctx.on('comment-deleted', (comment) => {
      const node = this.commentNodes.find(c => c.getID() === comment.id)
      if (!node) { console.error(`comment node id=${comment.id} not found`);return }
      node.getEl().remove()
      this.commentNodes = this.commentNodes.filter(c => c.getID() !== comment.id)
      // TODO: remove child nodes
    })

    // When comment update
    this.ctx.on('comment-updated', (comment) => {
      const node = this.commentNodes.find(c => c.getID() === comment.id)
      node && node.setData(comment)
    })
  }
}
