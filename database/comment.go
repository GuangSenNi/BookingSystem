package database

import "fmt"

func CustomAddComments(comment Comments) error {
	err = db.Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func QueryAllComments(fid int64) ([]Comments,error) {
	var comments []Comments
	err = db.Where("fid = ?",fid).Find(&comments).Error
	if err != nil {
		return comments, err
	}
	return comments, nil
}

func DelComments(phoneNumber, cTime string) bool {
	var comment Comments

	err = db.Where("phone_number=? AND ctime = ?",phoneNumber,cTime).First(&comment).Error
	if err != nil||comment.Cid==0 {
		fmt.Printf("cid=%v, DelComments err = %v",comment.Cid,err)
		return false
	}

	err = db.Where("cid=?",comment.Cid).Delete(Comments{}).Error
	if err != nil {
		fmt.Printf("DelComments err = %v",err)
		return false
	}
	return true
}
