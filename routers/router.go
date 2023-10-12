package routers

import (
	"mbook/controllers"

	beego "github.com/beego/beego"
)

func init() {
	//     beego.Router("/", &controllers.MainController{})

	// 首页 & 分类 & 详情
	beego.Router("/", &controllers.HomeController{}, "get:Index")               // 首页
	beego.Router("/explore", &controllers.ExploreController{}, "get:Index")     // 分类
	beego.Router("/books/:key", &controllers.DocumentController{}, "get:Index") // 分类

	// 读书
	beego.Router("/read/:key/:id", &controllers.DocumentController{}, "*:Read")         // 读图书章节
	beego.Router("/read/:key/search", &controllers.DocumentController{}, "post:Search") // 图书章节内搜索

	// 图书编辑
	beego.Router("/api/:key/edit/?:id", &controllers.DocumentController{}, "*:Edit")       // 图书编辑
	beego.Router("/api/:key/content/?:id", &controllers.DocumentController{}, "*:Content") // 图书内容获取
	beego.Router("/api/upload", &controllers.DocumentController{}, "post:Upload")          // 图书图片上传
	beego.Router("/api/:key/create", &controllers.DocumentController{}, "post:Create")     // 图书创建
	beego.Router("/api/:key/delete", &controllers.DocumentController{}, "post:Delete")     // 图书删除

	// 搜索
	beego.Router("/search", &controllers.SearchController{}, "get:Search")        // 搜索
	beego.Router("/search/result", &controllers.SearchController{}, "get:Result") // 搜索结果

	// 登录
	beego.Router("/login", &controllers.AccountController{}, "*:Login")              // 登录
	beego.Router("/register", &controllers.AccountController{}, "*:Register")        // 注册
	beego.Router("/logout", &controllers.AccountController{}, "*:Logout")            // 退登
	beego.Router("/doRegister", &controllers.AccountController{}, "post:DoRegister") // 完成注册

	// 用户图书管理
	beego.Router("/book", &controllers.BookController{}, "*:Index")                         // 我的图书
	beego.Router("/book/create", &controllers.BookController{}, "post:Create")              // 创建图书
	beego.Router("/book/:key/setting", &controllers.BookController{}, "*:Setting")          // 图书设置
	beego.Router("/book/setting/upload", &controllers.BookController{}, "post:UploadCover") // 图书封面设置
	beego.Router("/book/start/:id", &controllers.BookController{}, "*:Collection")          // 图书收藏
	beego.Router("/book/setting/save", &controllers.BookController{}, "post:Save")          // 图书保存
	beego.Router("/book/:key/release", &controllers.BookController{}, "post:Release")       // 图书发布
	beego.Router("/book/setting/token", &controllers.BookController{}, "post:CreateToken")  // 图书私有 Token 创建

	// 个人中心
	beego.Router("/user/:username", &controllers.UserController{}, "get:Index")                 // 分享
	beego.Router("/user/:username/collection", &controllers.UserController{}, "get:Collection") // 收藏
	beego.Router("/user/:username/follow", &controllers.UserController{}, "get:Follow")         // 关注
	beego.Router("/user/:username/fans", &controllers.UserController{}, "get:Fans")             // 粉丝
	beego.Router("/follow/:uid", &controllers.BaseController{}, "get:SetFollow")                // 关注或取消关注
	beego.Router("/book/score/:id", &controllers.BookController{}, "*:Score")                   // 评分
	beego.Router("/book/comment/:id", &controllers.BookController{}, "post:Comment")            // 评论

	// 个人设置
	beego.Router("/setting", &controllers.SettingController{}, "*:Index")         // 查看或编辑个人设置
	beego.Router("/setting/upload", &controllers.SettingController{}, "*:Upload") // 个人头像设置

	// 管理后台
	beego.Router("/manager/category", &controllers.ManagerController{}, "post,get:Category")            // 管理分类
	beego.Router("/manager/update-cate", &controllers.ManagerController{}, "post,get:UpdateCategory")   // 提交分类
	beego.Router("/manager/del-cate", &controllers.ManagerController{}, "post,get:DeleteCategory")      // 删除分类
	beego.Router("/manager/icon-cate", &controllers.ManagerController{}, "post,get:UpdateCategoryIcon") // 更新分类的图标
}
