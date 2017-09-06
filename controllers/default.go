package controllers

import (
	"github.com/astaxie/beego"
	"os"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	ip := os.Getenv("HOST")

	data := map[string]string{
		"language": "golang",
		"host":     ip,
		"output":   "hello world",
	}
	c.Data["json"] = data
	c.ServeJSON()
}
