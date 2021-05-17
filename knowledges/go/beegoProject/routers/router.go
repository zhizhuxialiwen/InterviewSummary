package routers

import (
	"beegoProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//网址：localhost:8080/
	beego.Router("/", &controllers.MainController{})//访问默认的Get
	//网址：localhost:8080/goods
	beego.Router("/goods", &controllers.GoodsController{})//访问默认的Get
	
	//网址：localhost:8080/article
	beego.Router("/article", &controllers.ArticleController{}) //访问默认的Get方法
	beego.Router("/article/add", &controllers.ArticleController{}, "get:AddArticle")//访问自定义的Get
	beego.Router("/article/edit", &controllers.ArticleController{}, "get:EditArticle")
}
