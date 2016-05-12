package routers

import (
	"github.com/lomoalbert/gogt-sdk/demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/register",&controllers.RegisterController{})
	beego.Router("/validate",&controllers.ValidateController{})
	beego.Router("/ajax_validate",&controllers.AjaxValidateController{})
}
