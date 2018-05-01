package routers

import (
	"github.com/fedir/gmm02/beego-web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
