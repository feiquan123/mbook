package controllers

import (
	"mbook/models"

	"github.com/beego/beego"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	if categories, err := new(models.Category).GetCategories(-1, true); err != nil {
		c.Data["Cates"] = categories
	} else {
		beego.Error(err)
	}
	c.TplName = "home/list.html"
}
