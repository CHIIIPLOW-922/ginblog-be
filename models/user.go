package models

type User struct {
	ID       uint64 `json:"id"`
	Username string `gorm:"type:varchar(20);not null " json:"username,omitempty" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password,omitempty" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `json:"role"`
}

func (User) TableName() string {
	return "user"
}
