package action

import (
	"BookingSystem/database"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//上传图片 废弃
//func UploadPicture(c *gin.Context)  {
//	file, err := c.FormFile("upload")
//	if err != nil {
//		c.String(http.StatusBadRequest, "请求失败")
//		return
//	}
//	//获取文件名
//	fileName := file.Filename
//	fmt.Println("文件名：", fileName)
//	//保存文件到服务器本地
//	//SaveUploadedFile(文件头，保存路径)
//	uploadPictureUrl="assets/images/"+strconv.FormatInt(time.Now().Unix(), 10)+fileName
//	if err := c.SaveUploadedFile(file, uploadPictureUrl); err != nil {
//		fmt.Printf("SaveUploadedFile,err=%v", err)
//		c.JSON(http.StatusBadRequest, gin.H{
//			"message": err,
//		})
//		return
//	}
//
//}

//修改菜品 接收ajax formdata数据
func ChangeProduct(c *gin.Context)  {
	name := c.PostForm("name")
	price:= c.PostForm("price")
	detail := c.PostForm("detail")
	file, err := c.FormFile("file")
	var uploadPictureUrl string
	if err != nil {
		fmt.Printf("FormFile,err=%v", err)
	}else{
		//获取文件名
		fileName := file.Filename
		fmt.Println("文件名：", fileName)
		//保存文件到服务器本地
		//SaveUploadedFile(文件头，保存路径)
		uploadPictureUrl ="assets/images/"+strconv.FormatInt(time.Now().Unix()%1000, 10)+fileName
		if err := c.SaveUploadedFile(file, uploadPictureUrl); err != nil {
			fmt.Printf("SaveUploadedFile,err=%v", err)
			c.String(http.StatusBadRequest, "请求失败")
			return
		}
	}

	err = database.AlterFood(name,uploadPictureUrl,price,detail)
	if err!=nil{
		fmt.Printf("AlterFood fail,err=%v", err)
		return
	}
}
//新增菜品
func AddProduct(c *gin.Context)  {
	name := c.PostForm("name")
	price:= c.PostForm("price")
	kind:=c.PostForm("kind")
	detail := c.PostForm("detail")
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Printf("FormFile,err=%v", err)
		c.String(http.StatusBadRequest, "请求失败")
		return
	}
	//获取文件名
	fileName := file.Filename
	fmt.Println("文件名：", fileName)
	//保存文件到服务器本地
	//SaveUploadedFile(文件头，保存路径)
	uploadPictureUrl:="assets/images/"+strconv.FormatInt(time.Now().Unix()%1000, 10)+fileName
	if err := c.SaveUploadedFile(file, uploadPictureUrl); err != nil {
		fmt.Printf("SaveUploadedFile,err=%v", err)
		c.String(http.StatusBadRequest, "请求失败")
		return
	}
	err = database.AddFood(name,uploadPictureUrl,price,kind,detail)
	if err!=nil{
		fmt.Printf("AddFood fail,err=%v", err)
		return
	}

}
//查询所有未完成订单
func QueryAllUserOrders(c *gin.Context)  {
	cmd := c.Param("cmd")
	orderList,err := database.QueryUserOrders(cmd)
	if err!=nil{
		fmt.Printf("QueryOrders fail,err=%v", err)
		return
	}
	c.JSON(http.StatusOK, orderList)
}
//修改订单状态
func ChangeOrder(c *gin.Context)  {
	body, _ := ioutil.ReadAll(c.Request.Body)
	data := make(map[string]string)
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("ChangeOrder Unmarshal fail,err=%v", err)
		return
	}
	err=database.AlterOrder(data["name"],data["time"],data["state"])
	if err!=nil{
		fmt.Printf("AlterOrder fail,err=%v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}