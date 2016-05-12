package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lomoalbert"
)

type MainController struct {
	beego.Controller
}

type RegisterController struct {
	beego.Controller
}

type ValidateController struct {
	beego.Controller
}

type AjaxValidateController struct {
	beego.Controller
}

func (ctl *MainController) Get() {
	ctl.TplName = "index.tpl"
}

func (c *RegisterController)Get(){
	user_id:="test"

}