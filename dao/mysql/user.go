package mysql

import "ginblog-be/models"

func GetAllUsers() (users []models.User, total int64) {
	db.Select("id,username,role").Find(&users)
	db.Model(&users).Count(&total)
	if db.Error != nil {
		return users, total
	}
	return users, total
}
