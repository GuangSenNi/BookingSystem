package main


import (
	"BookingSystem/database"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(111)
	database.InitDB()

	r := gin.Default()
	//路径映射
	r.GET("/user/init", InitPage)
	r.POST("/user/create", CreateUser)

	//端口号
	r.Run(":8080")
}
func InitPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func CreateUser(c *gin.Context) {

}