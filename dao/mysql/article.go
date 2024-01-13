package mysql

import "ginblog-be/models"

func GetAllArticle() (articles []models.Article, total int64) {
	db.Select("*").Find(&articles).Count(&total)
	//db.Model(&articles).Count(&total)
	if db.Error != nil {
		return articles, total
	}
	return articles, total
}
