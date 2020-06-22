package database

import "fmt"

func CustomLogin(phoneNumber, password string) bool {
	var user UserInfos
	err := db.Where("phone_number = ? AND password = ?", phoneNumber, password).First(&user).Error
	if err != nil {
		fmt.Printf("CustomLogin fail, err=%v \n", err)
		return false
	}
	return true
}

func CreateUser(phoneNumber, username, password string) error {
	var user UserInfos
	err := db.Where("phone_number = ?", phoneNumber).First(&user).Error
	if err == nil {
		fmt.Printf("CustomLogin fail, err=%v \n", err)
		return fmt.Errorf("phone_number must be unique")
	}
	user = UserInfos{
		UserName:    username,
		PhoneNumber: phoneNumber,
		Password:    password,
	}
	err = db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
