package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type ApiController struct {
	beego.Controller
}

type MotController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func(c *ApiController) Get() {
	c.TplName = "api.tpl"
}

func(c *MotController) Get() {
	c.TplName = "mot.tpl"
}
