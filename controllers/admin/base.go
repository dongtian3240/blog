package admin

import (
	"blog/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	UserId   int
	UserName string
}

func (this *BaseController) Prepare() {

	fmt.Println("=======Prepare======")
	var m map[string]string = make(map[string]string)
	m["Css"] = "admin/css.html"
	m["HtmlHead"] = "admin/header.html"
	m["SideBar"] = "admin/sidebar.html"
	m["Scripts"] = "admin/scripts.html"
	this.LayoutSections = m
	this.Layout = "admin/layout.html"

	CheckLogin(this)

}

func CheckLogin(this *BaseController) {

	userId := this.Ctx.Input.Query(":userId")
	fmt.Println("====userId", userId)
	user := this.GetSession("user")
	//没有登录
	if user == nil {
		this.Redirect("/login", http.StatusFound)
		return
	} else { //已经登录判断是否是当前用户
		if u, ok := user.(models.User); ok {

			uId, err := strconv.Atoi(userId)
			if err != nil {
				this.Abort("404")
				return
			} else if uId != u.UserId {
				this.Abort("401")
			} else if u.UserId == uId {
				this.Data["user"] = u
				this.Data["userId"] = userId
				this.UserId = u.UserId
				this.UserName = u.UserName
			} else {
				this.Abort("404")
				return
			}

		} else {

			this.Abort("401")
			return
		}
	}
}
