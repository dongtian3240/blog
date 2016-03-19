package admin

type CommentController struct {
	BaseController
}

func (this *CommentController) Index() {

	this.TplName = "admin/comment/comment.html"
}
