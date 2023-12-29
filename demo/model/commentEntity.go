package model

type Comment struct {
	Id           uint   `gorm:"primary_key" json:"id"`
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

func GetAllComments() ([]Comment, int64) {
	var comments []Comment
	var total int64
	db.Select("*").Find(&comments)
	db.Model(&comments).Count(&total)
	if err := db.Error; err != nil {
		return comments, 0
	}
	return comments, total
}
