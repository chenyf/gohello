package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	data := map[string]string{
		"language": "golang",
		"output":   "hello world",
	}
	c.Data["json"] = data
	c.ServeJSON()
}
