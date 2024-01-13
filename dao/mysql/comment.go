package mysql

import "ginblog-be/models"

func GetAllComments() (comments []models.Comment, total int64) {
	db.Select("*").Find(&comments).Count(&total)
	if db.Error != nil {
		return comments, total
	}
	return comments, total
}
