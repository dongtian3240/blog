package admin

type AdminController struct {
	BaseController
}

func (this *AdminController) Index() {

	this.TplName = "admin/index.html"
}
