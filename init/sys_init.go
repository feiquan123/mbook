package init

import (
	"path/filepath"

	"github.com/beego/beego"
)

func sysInit() {
	// 指定前端上传路径
	uploadDir := filepath.Join("./", "uploads")
	beego.BConfig.WebConfig.StaticDir["uploads"] = uploadDir

	// 注册前端函数
	registerFuncs()
}

func registerFuncs() {
	beego.AddFuncMap("cdnjs", func(path string) string {
		cdnJSPath := beego.AppConfig.DefaultString("cdnjs", "")
		return filepath.Join(cdnJSPath, path)
	})
}
