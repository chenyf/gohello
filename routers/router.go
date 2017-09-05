package routers

import (
	"github.com/chenyf/gohello/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
