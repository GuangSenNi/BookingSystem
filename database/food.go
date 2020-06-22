package database

import "fmt"

//增加数量
func AddCount(name string,count int64) error {
	var food Foods

	err = db.Where("food_name = ? ", name).First(&food).Error
	if err != nil {
		return  err
	}
	row:=db.Model(Foods{}).Where("fid=?",food.Fid).UpdateColumn(Foods{Count: food.Count+count}).RowsAffected
	fmt.Println("###AlterFood 影响的行",row)
	return nil
}

//添加收藏
func AddFavorite(fid int64,phoneNumber string) error {
	fav := FavoriteFood{
		PhoneNumber:    phoneNumber,
		Fid: fid,
	}
	err = db.Create(&fav).Error
	if err != nil {
		return err
	}
	return nil
}
//删除收藏
func DeleteFavorite(fid int64,phoneNumber string) error {
	fav := FavoriteFood{
		PhoneNumber:    phoneNumber,
		Fid: fid,
	}
	err = db.Delete(&fav).Error
	if err != nil {
		return err
	}
	return nil
}

//查询所有菜品
func QueryAllFood() ([]Foods, error) {
	var foods []Foods
	err = db.Find(&foods).Error
	if err != nil {
		return foods, err
	}
	return foods, nil
}

//根据种类查菜品
func QueryByKinds(kinds string) ([]Foods, error) {
	var foods []Foods
	err = db.Where("food_kinds = ? ", kinds).Find(&foods).Error
	if err != nil {
		return foods, err
	}
	return foods, nil
}

//查询收藏的菜品 可能有错 重点测试
func QueryByFav(phoneNumber string) ([]Foods, error) {
	var foods []Foods
	query := "SELECT foods.fid,foods.food_name,foods.price,foods.picture,foods.brief,foods.food_kinds " +
		"FROM foods,favorite_food where favorite_food.fid=foods.fid and favorite_food.phone_number= ? "
	//执行原生语句
	err = db.Raw(query, phoneNumber).Scan(&foods).Error
	if err != nil {
		return foods, err
	}
	return foods, nil
}
