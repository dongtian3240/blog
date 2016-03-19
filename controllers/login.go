package controllers

import (
	"blog/common"
	"blog/models"
	"net/http"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

/****
跳转
*/
func (this *LoginController) Login() {

	user := this.GetSession("user")
	if user != nil {

		this.Redirect("/", http.StatusFound)

	} else {
		this.TplName = "login.tpl"
	}

}

/****
用户登录
***/
func (this *LoginController) DoLogin() {

	result := common.JsonResult{}
	userName := this.GetString("userName", "")
	password := this.GetString("password", "")
	//remember := this.GetBool("remember", false)

	if userName == "" {

		result.Message = "用户名不能为空"
	} else if password == "" {
		result.Message = "密码不能为空"
	}

	u, err := models.Login(userName, password)

	if err != nil {
		result.Message = "用户名或密码错误"
	} else {
		this.SetSession("user", u)
		result.Message = "登录成功！"
		result.Success = true
		result.Data = u
	}

	this.Ctx.Output.JSON(result, true, false)

}

/***
注销登录
*/
func (this *LoginController) Logout() {

	this.DelSession("user")
	this.Redirect("/", http.StatusFound)
}
