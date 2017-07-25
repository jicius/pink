package routers

import (
	"pink/controllers"
	"github.com/astaxie/beego"
)

// 路由注册，将URL映射到Controller
func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api", &controllers.ApiController{})
	beego.Router("/mot", &controllers.MotController{})
}
