package main

import (
	_ "mbook/init"
	_ "mbook/routers"

	"github.com/beego/beego"
)

func main() {
	beego.Run()
}
