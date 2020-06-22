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
//下单 order_success下单成功 order_fail下单失败 wait2eat制作完成 finish消费结束
//为简化状态存中文名
//todo 调用支付宝接口
func CustomAddOrder(c *gin.Context)  {
	session := sessions.Default(c)
	phoneNumber := session.Get("username").(string)
	body, _ := ioutil.ReadAll(c.Request.Body)
	data := make(map[string]string)
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("CustomAddOrder Unmarshal fail,err=%v", err)
		return
	}
	price,_:=strconv.ParseFloat(data["sumPrice"], 64)
	code:=strconv.FormatInt(time.Now().Unix()%10000,10)
	order:=database.Orders{
		PhoneNumber: phoneNumber,
		Price:       price,
		State:       "下单成功",
		StartTime:   time.Now().Format("2006-01-02 15:04:05"),
		Detail:      data["detail"],
		MealTime:data["mealTime"],
		MealCode:	 code,
	}
	err = database.AddOrders(order)
	if err!=nil{
		fmt.Printf("CustomAddOrder AddOrders fail,err=%v", err)
		return
	}
	//增加count
	data2 := make(map[string]string)
	detail :=data["detail"]
	err = json.Unmarshal([]byte(detail), &data2)
	if err != nil {
		fmt.Printf("detail Unmarshal fail,err=%v", err)
		return
	}
	fmt.Printf("data2:%v \n",  data2)
	for k, v := range data2 {
		count, _ := strconv.ParseInt(v, 10, 64)
		err=database.AddCount(k,count)
		if err != nil {
			fmt.Printf("AddCount fail,err=%v", err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func QueryOrders(c *gin.Context)  {
	cmd := c.Param("cmd")
	session := sessions.Default(c)
	phoneNumber := session.Get("username").(string)
	orderList,err := database.QueryOrders(cmd,phoneNumber)
	if err!=nil{
		fmt.Printf("QueryOrders fail,err=%v", err)
		return
	}
	c.JSON(http.StatusOK, orderList)
}