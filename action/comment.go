package action

import (
	"BookingSystem/database"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)
//添加评论
func AddComment(c *gin.Context) {
	session := sessions.Default(c)
	phoneNumber := session.Get("username").(string)
	body, _ := ioutil.ReadAll(c.Request.Body)
	data := make(map[string]string)
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("AddComment Unmarshal fail,err=%v", err)
		return
	}
	fid, _ := strconv.ParseInt(data["fid"], 10, 64)
	comment:=database.Comments{
		PhoneNumber:phoneNumber,
		Ctime:time.Now().Format("2006-01-02 15:04:05"),
		Fid:fid,
		Message:data["message"],
	}
	err =database.CustomAddComments(comment)
	if err!=nil{
		fmt.Printf("CustomAddComments err, err=%v",err)
		return
	}
	comments,_:=database.QueryAllComments(fid)
	c.JSON(http.StatusOK, comments)

}
//查看评论
func QueryComments(c *gin.Context)  {
	fidStr := c.Param("fid")
	fid, _ := strconv.ParseInt(fidStr, 10, 64)
	comments,err:=database.QueryAllComments(fid)
	if err!=nil{
		fmt.Printf("QueryAllComments err, err=%v",err)
		return
	}
	c.JSON(http.StatusOK, comments)
}
//删除评论
func DelComments(c *gin.Context)  {
	session := sessions.Default(c)
	phoneNumber := session.Get("username").(string)
	body, _ := ioutil.ReadAll(c.Request.Body)
	data := make(map[string]string)
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("DelComments Unmarshal fail,err=%v", err)
		return
	}
	database.DelComments(phoneNumber,data["ctime"])
	fid, _ := strconv.ParseInt(data["fid"], 10, 64)
	comments,_:=database.QueryAllComments(fid)
	c.JSON(http.StatusOK, comments)
}