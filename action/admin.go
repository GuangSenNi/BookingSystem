package action

import (
	"BookingSystem/database"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Login(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	data := make(map[string]string)
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("login Unmarshal fail,err=%v", err)
		return
	}
	ok := database.CustomLogin(data["username"], data["password"])
	if ok {
		//设置cookie 生存期的单位是秒
		c.SetCookie("user_cookie", data["username"], 1000, "/", "localhost", false, true)
		//设置session
		session := sessions.Default(c)
		session.Set("username", data["username"])
		session.Save()

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "invalid username or password",
	})
}
//管理员登录
func Login2(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	data := make(map[string]string)
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("login Unmarshal fail,err=%v", err)
		return
	}
	if data["password"]=="admin" {
		//设置cookie 生存期的单位是秒
		c.SetCookie("user_cookie", "admin", 1000, "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "invalid username or password",
	})
}

func Register(c *gin.Context) {
	data := make(map[string]string)
	body, _ := ioutil.ReadAll(c.Request.Body)
	println("body=" + string(body))
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Register Unmarshal fail,err=%v \n", err)
		return
	}
	err = database.CreateUser(data["phoneNumber"], data["username"], data["password"])
	if err != nil {
		fmt.Printf("CreateUser fail,err=%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
