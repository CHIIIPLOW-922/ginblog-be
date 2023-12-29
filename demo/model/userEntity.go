package model

type User struct {
	ID       uint64 //`json:"id"`
	Username string //`gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string //`gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    //`gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

func (User) TableName() string {
	return "user"
}

func GetUserById(id int) User {
	var user User
	db.Select("id,username,role").Where("id = ?", id).First(&user)
	if db.Error != nil || user.ID == 0 || user == (User{}) {
		return User{}
	}
	return user
}

func GetAllUsers() ([]User, int64) {
	var users []User
	var total int64
	db.Select("id,username,role").Find(&users)
	db.Model(&users).Count(&total)
	if err != nil {
		return users, 0
	}
	return users, total
}
