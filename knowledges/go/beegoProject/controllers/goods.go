package controllers
import (
	"github.com/astaxie/beego"
)


type GoodsController struct {
	beego.Controller
}

//控制器： 业务逻辑
func (c *GoodsController) Get() {
	//加载view/index.tpl
	c.Data["title"] = "你好golang" //绑定数据
	c.Data["num"] = 12
	c.TplName = "goods.tpl"
}