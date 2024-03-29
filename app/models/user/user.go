package user

import (
	"goblog/app/models"
	"goblog/pkg/password"
	"goblog/pkg/route"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `gorm:"column:name;type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"column:email;type:varchar(255);default:NULL;unique;" valid:"email"`
	Password string `gorm:"column:password;type:varchar(255)" valid:"password"`
	// gorm:"-" —— 设置 GORM 在读写时略过此字段
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

// ComparePassword 对比密码是否匹配
func (u User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, u.Password)
}

// Link 方法用来生成用户链接
func (u User) Link() string {
	return route.Name2URL("users.show", "id", u.GetStringID())
}
