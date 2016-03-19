package main

import (
	_ "blog/routers"

	"blog/functions"

	"github.com/astaxie/beego"
)

func main() {
	//设置静态文件的路径
	beego.SetStaticPath("/static", "static")
	beego.SetStaticPath("/public", "public")

	//设置模板
	beego.SetViewsPath("views")
	beego.AddFuncMap("increment", functions.Increment)
	beego.Run()
}
