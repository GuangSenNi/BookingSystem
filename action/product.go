package action

import (
	"BookingSystem/database"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
//添加收藏
func Add2Favorite(c *gin.Context)  {
	fidStr := c.Param("fid")
	fmt.Println("fidStr=",fidStr)
	session := sessions.Default(c)
	phoneNumber := session.Get("username").(string)
	fid, _ := strconv.ParseInt(fidStr, 10, 64)
	err:=database.AddFavorite(fid,phoneNumber)
	if err!=nil{
		fmt.Printf("Add2Favorite fail ,err=%v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
//删除收藏
func DelFav(c *gin.Context)  {
	fidStr := c.Param("fid")
	session := sessions.Default(c)
	phoneNumber := session.Get("username").(string)
	fid, _ := strconv.ParseInt(fidStr, 10, 64)
	err:=database.DeleteFavorite(fid,phoneNumber)
	if err!=nil{
		fmt.Printf("DelFav fail ,err=%v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func FindByKind(c *gin.Context) {
	//普通盖码 gaiMa、特色小吃 xiaoChi、精致小炒 xiaoChao、水饺汤面 tangMian , 收藏 favorite
	kinds := c.Param("kinds")
	fmt.Println("kinds=",kinds)
	if kinds == "favorite" {
		session := sessions.Default(c)
		phoneNumber := session.Get("username").(string)
		fmt.Println("phone number=", phoneNumber)
		foods, err := database.QueryByFav(phoneNumber)
		if err != nil {
			fmt.Printf("QueryByFav fail ,err=%v", err)
			return
		}
		c.JSON(http.StatusOK, foods)
	} else if kinds == "all"{
		foods, err := database.QueryAllFood()
		if err != nil {
			fmt.Printf("FindAllProducts fail ,err=%v", err)
			return
		}
		c.JSON(http.StatusOK, foods)
	}else {
		foods, err := database.QueryByKinds(kinds)
		if err != nil {
			fmt.Printf("QueryByKinds fail ,err=%v", err)
			return
		}
		c.JSON(http.StatusOK, foods)
	}
}

func FindAllProducts(c *gin.Context) {
	foods, err := database.QueryAllFood()
	if err != nil {
		fmt.Printf("FindAllProducts fail ,err=%v", err)
		return
	}
	c.JSON(http.StatusOK, foods)
}

func ProductList(c *gin.Context) {
	c.HTML(http.StatusOK, "product-grid.html", gin.H{
		"title": "index",
	})
}
