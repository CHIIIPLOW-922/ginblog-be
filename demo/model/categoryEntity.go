package model

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func (Category) TableName() string {
	return "category"
}

func GetAllCategories() ([]Category, int64) {
	var categories []Category
	var total int64
	db.Select("*").Find(&categories).Count(&total)
	if len(categories) == 0 || db.Error != nil {
		return categories, 0
	}
	return categories, total
}
