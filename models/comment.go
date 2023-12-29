package models

type Comment struct {
	ID           uint   `gorm:"primary_key" json:"id"`
	UserId       uint   `json:"user_id"`
	ArticleId    uint   `json:"article_id"`
	ArticleTitle string `gorm:"type:varchar(500);not null;" json:"article_title"`
	Username     string `gorm:"type:varchar(500);not null;" json:"username"`
	Content      string `gorm:"type:varchar(500);not null;" json:"content"`
	Status       int8   `gorm:"type:tinyint;default:2" json:"status"`
}

func (Comment) TableName() string {
	return "comment"
}
