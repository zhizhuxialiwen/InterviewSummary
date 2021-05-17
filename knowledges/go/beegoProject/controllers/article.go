package controllers
import (
	"github.com/astaxie/beego"
	"fmt"
)


type ArticleController struct {
	beego.Controller
}

//控制器： 业务逻辑
func (c *ArticleController) Get() {
	c.Ctx.WriteString("新闻列表") //直接给页面返回数据，不需要连接view
}

//新增方法1
func (c *ArticleController) AddArticle() {
	c.Ctx.WriteString("增加新闻") //直接给页面返回数据，不需要连接view
}
//新增方法2
func (c *ArticleController) EditArticle() {
	//获取get传值
	//id := c.GetString("id")
	//beego.Info(id)
	//fmt.Printf("值：%v, 类型：%v \n", id, id)
	//c.Ctx.WriteString("编辑新闻--" + id) //直接给页面返回数据，不需要连接view

	id,err := c.GetInt("id")
	if err != nil {
		beego.Info(id)
		c.Ctx.WriteString("错误编辑新闻--") 
	}
	fmt.Printf("值：%v, 类型：%v \n", id, id)
	c.Ctx.WriteString("编辑新闻--") 
}