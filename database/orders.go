package database
//新建订单
func AddOrders(order Orders) error {
	err = db.Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}
//查询订单
func QueryOrders(cmd,phoneNumber string) ([]Orders,error) {
	var orderList []Orders
	if cmd=="unfinished"{
		err = db.Where("(state = ? or state = ? ) AND phone_number = ?", "下单成功","制作完成",phoneNumber).Find(&orderList).Error
	}else{
		err = db.Where("phone_number = ?",phoneNumber).Find(&orderList).Error
	}
	if err != nil {
		return orderList, err
	}
	return orderList,nil
}
//修改状态

//修改取餐码和窗口



