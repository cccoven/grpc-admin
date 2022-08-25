package model

// UserRole 用户与角色关联表
type UserRole struct {
	UserID    uint `gorm:"user_id"`
	RoleID    uint `gorm:"role_id"`
}

func (u UserRole) TableName() string {
	return "ga_user_roles"
}
