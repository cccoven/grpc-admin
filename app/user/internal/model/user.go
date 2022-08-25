package model

import "grpc-admin/common/pkg"

type User struct {
	pkg.GormModel
	Username string `gorm:"comment:用户名" json:"username"`
	Password string `gorm:"comment:密码" json:"password"`
	Gender   int    `gorm:"type:tinyint" json:"gender"`
	Avatar   string `gorm:"comment:头像" json:"avatar"`
	Phone    string `gorm:"comment:手机号码" json:"phone"`
	Roles    []Role `gorm:"many2many:ga_user_roles" json:"-"` // 用户角色（多对多）
}

func (u User) TableName() string {
	return "ga_users"
}
