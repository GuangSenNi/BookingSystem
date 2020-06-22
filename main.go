package main

import (
	"BookingSystem/action"
	"BookingSystem/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	database.InitDB()

	r := gin.Default()
	//使用session
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	//加载静态资源
	r.Static("custom/assets", "./assets")
	//加载HTML模板
	r.LoadHTMLGlob("template/*")
	//路径映射
	//客户端
	r.GET("/init", InitPage)
	r.POST("/login", action.Login)
	r.POST("/register", action.Register)
	//管理员登录
	r.GET("/adminInit", AdminInitPage)
	r.POST("/adminLogin", action.Login2)

	custom := r.Group("custom",AuthMiddleware)
	{
		custom.GET("/index", Index)
		//查找所有菜品
		//custom.POST("/productList", action.FindAllProducts)
		//跳转到菜品页
		custom.GET("/product-grid.html", action.ProductList)
		//根据种类查询菜品
		custom.POST("/searchByKinds/:kinds", action.FindByKind)
		//添加收藏
		custom.POST("/add2fav/:fid", action.Add2Favorite)
		//删除收藏
		custom.POST("/delFav/:fid", action.DelFav)
		//下单
		custom.POST("/sendOrder", action.CustomAddOrder)
		//查看所有评论
		custom.POST("/queryComments/:fid", action.QueryComments)
		//添加评论
		custom.POST("/addComments", action.AddComment)
		//删除评论
		custom.POST("/delComments", action.DelComments)
		//查看订单
		custom.POST("/queryOrders/:cmd", action.QueryOrders)

		//管理员端功能
		//管理端首页
		custom.GET("/product-grid2.html", AdminIndex)
		//修改菜
		custom.POST("/changeProduct", action.ChangeProduct)
		//上传菜图片不修改 废弃
		//custom.POST("/changePic", action.UploadPicture)
		//新增菜
		custom.POST("/addFood", action.AddProduct)
		//查询所有用户订单
		custom.POST("/queryOrders2/:cmd", action.QueryAllUserOrders)
		//修改订单状态
		custom.POST("/changeOrder", action.ChangeOrder)

	}

	//端口号
	r.Run(":8080")
}
func AdminInitPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login2.html", gin.H{
		"title": "login",
	})
}
func AdminIndex(c *gin.Context)  {
	c.HTML(http.StatusOK, "product-grid2.html", gin.H{
		"title": "index",
	})
}
func InitPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "login",
	})
}
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "index",
	})
}

//权限验证
//todo 重定向有问题
func AuthMiddleware(c *gin.Context) {
	//c.Next()
	_, e := c.Request.Cookie("user_cookie")
	if e == nil {
		c.Next()
	} else {
		c.Abort()
		//c.Redirect(http.StatusMovedPermanently,"http://localhost:8080/init")
		c.String(http.StatusForbidden, "need login")
	}
}
