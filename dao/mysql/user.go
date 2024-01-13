package mysql

import (
	"errors"
	"ginblog-be/models"
	"ginblog-be/utils/snowflake"
)

func GetAllUsers() (users []models.User, total int64) {
	db.Select("id,username,role").Find(&users).Count(&total)
	//db.Model(&users).Count(&total)
	if db.Error != nil {
		return users, total
	}
	return users, total
}

func InsertUser(user *models.User) (err error) {
	user.ID, err = snowflake.GetID()
	if err != nil {
		return err
	}
	db.Create(&user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func IsDuplicated(user *models.User) (err error) {
	var num int64
	db.Model(&user).Where("username", &user.Username).Count(&num)
	if db.Error != nil {
		return db.Error
	}
	if num > 0 {
		return errors.New("用户已存在")
	}
	return nil
}
