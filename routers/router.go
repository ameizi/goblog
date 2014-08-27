package routers

import (
	"github.com/astaxie/beego"
	"goblog/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})

	beego.Router("/captcha", &controllers.UserController{}, "get:NewCaptcha") //验证码

	beego.Router("/login", &controllers.UserController{}, "get:Login") //用户登录页

	beego.Router("/login", &controllers.UserController{}, "post:LoginAction") //用户登录Action

	beego.Router("/logout", &controllers.UserController{}, "get:Logout") //用户退出

	beego.Router("/register", &controllers.UserController{}, "get:Register") //用户注册页面

	beego.Router("/register", &controllers.UserController{}, "post:RegisterAction") //用户注册Action

	beego.Router("/user/delete/:id:int", &controllers.UserController{}, "get:DeleteUser") //删除用户

	beego.Router("/user/view/:username:string", &controllers.UserController{}, "get:GetUserInfo") //获取用户详情

	beego.Router("/topic", &controllers.TopicController{})

	beego.Router("/topic/add", &controllers.TopicController{}, "post:AddTopic") //发表文章

	beego.Router("/topic/view/:id:int", &controllers.TopicController{}, "get:ViewTopic") //查看文章

	beego.Router("/topic/delete/:id:int", &controllers.TopicController{}, "get:DeleteTopic") //删除文章

	beego.AutoRouter(&controllers.TopicController{})

	beego.Router("/category", &controllers.CategoryController{}, "get:GetAllCategory") //获取所有的分类

	beego.Router("/category/add", &controllers.CategoryController{}, "post:AddCategory") //添加文章分类

	beego.Router("/category/modify", &controllers.CategoryController{}, "post:ModifyCategory") //修改文章分类

	beego.Router("/category/delete/:id:int", &controllers.CategoryController{}, "get:DeleteCategory") //删除文章分类

	beego.AutoRouter(&controllers.CategoryController{})

}
