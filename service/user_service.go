package service

import (
	"ginblog-be/dao/mysql"
	"ginblog-be/models"
	"ginblog-be/utils/snowflake"
)

func SaveUser(user *models.User) error {
	if err := mysql.IsDuplicated(user); err != nil {
		return err
	}
	user.ID, _ = snowflake.GetID()
	if err := mysql.InsertUser(user); err != nil {
		return err
	}
	return nil

}
