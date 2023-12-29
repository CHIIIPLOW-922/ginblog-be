package mysql

import (
	"ginblog-be/models"
	"ginblog-be/utils/snowflake"
)

func GetAllUsers() (users []models.User, total int64) {
	db.Select("id,username,role").Find(&users)
	db.Model(&users).Count(&total)
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
