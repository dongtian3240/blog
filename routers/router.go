package routers

import (
	"blog/controllers"
	"blog/controllers/admin"
	"blog/upload"

	"github.com/astaxie/beego"
)

func init() {

	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.IndexController{}, "get:Index")

	beego.Router("/login", &controllers.LoginController{}, "get:Login")
	beego.Router("/login", &controllers.LoginController{}, "post:DoLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	beego.Router("/admin/:userId", &admin.AdminController{}, "*:Index")

	beego.Router("/admin/:userId/topic", &admin.TopicController{}, "*:Index")
	beego.Router("/admin/:userId/topic/new", &admin.TopicController{}, "get:ToNewTopic")
	beego.Router("/admin/:userId/topic/new", &admin.TopicController{}, "post:NewTopic")
	beego.Router("/admin/:userId/topic/edit", &admin.TopicController{}, "get:ToEditTopic")
	beego.Router("/admin/:userId/topic/edit", &admin.TopicController{}, "post:EditTopic")
	beego.Router("/admin/:userId/topic/deleteTopic", &admin.TopicController{}, "post:DeleteTopic")

	beego.Router("/admin/:userId/label", &admin.LabelController{}, "*:Index")
	beego.Router("/admin/:userId/label/new", &admin.LabelController{}, "post:NewLabel")
	beego.Router("/admin/:userId/label/findLabelById", &admin.LabelController{}, "post:FindLabelById")
	beego.Router("/admin/:userId/label/update", &admin.LabelController{}, "post:UpdateLabel")
	beego.Router("/admin/:userId/label/delete", &admin.LabelController{}, "post:DeleteLabelById")

	beego.Router("/admin/:userId/label/findLabelListByTopicId", &admin.LabelController{}, "post:FindLabelListByTopicId")
	beego.Router("/admin/:userId/comment", &admin.CommentController{}, "*:Index")

	beego.Router("/upload/uploadImage", &upload.UploadController{}, "post:UploadImage")
}
