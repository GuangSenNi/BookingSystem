package database

import (
	"fmt"
	"strconv"
	"time"
)
//修改菜品
func AlterFood(name,picUrl,price,info string) error {
	var food Foods
	var p float64
	err = db.Where("food_name = ? ", name).First(&food).Error
	if err != nil {
		return  err
	}
	if price!=""{
		p, _ = strconv.ParseFloat(price, 64)
	}

	row:=db.Model(Foods{}).Where("fid=?",food.Fid).UpdateColumn(Foods{Picture: picUrl, Price: p,Brief:info}).RowsAffected
	fmt.Println("###AlterFood 影响的行",row)

	return nil
}
//新增菜品
func AddFood(name,picUrl,price,kind,info string) error {
	var food Foods
	err = db.Where("food_name = ? ", name).First(&food).Error
	if err == nil {
		return  fmt.Errorf("name must unique")
	}
	if picUrl==""{
		return  fmt.Errorf("picUrl is null")
	}
	if price==""{
		return  fmt.Errorf("price is null")

	}
	if kind ==""{
		return  fmt.Errorf("kind is null")
	}
	v, _ := strconv.ParseFloat(price, 64)
	food=Foods{
		Fid:time.Now().Unix()%10000,
		FoodName:name,
		Price:v,
		Picture:picUrl,
		Brief:info,
		FoodKinds:kind,
	}
	err = db.Create(&food).Error
	if err != nil {
		return err
	}
	return nil
}
//管理员查询未完成订单
func QueryUserOrders(cmd string) ([]Orders,error) {
	var orderList []Orders
	if cmd=="unfinished"{
		err = db.Where("(state = ? or state = ? )", "下单成功","制作完成").Find(&orderList).Error
	}else{
		err = db.Find(&orderList).Error
	}
	if err != nil {
		return orderList, err
	}
	return orderList,nil
}
//修改订单状态
func AlterOrder(phoneNumber,startTime,state string) error {
	row:=db.Model(Orders{}).Where("phone_number = ? AND start_time = ?", phoneNumber,startTime).UpdateColumn("state", state).RowsAffected
	fmt.Println("#####rows=",row)
	if row == 0 {
		return  fmt.Errorf("0 row affect")
	}
	return nil
}