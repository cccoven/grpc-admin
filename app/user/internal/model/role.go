package model

import (
	"grpc-admin/common/pkg"
	"time"
)

type Role struct {
	pkg.GormModel
	DeletedAt *time.Time `json:"-"`
	Name      string     `gorm:"comment:角色名称" json:"name"`                                       // 角色名称
	Desc      string     `gorm:"comment:角色描述" json:"desc"`                                       // 角色描述
	ParentID  uint       `gorm:"column:parent_id;comment:父角色 id" json:"parentID"`                // 父角色id
	IsDefault int        `gorm:"type:tinyint;default:(0);comment:是否为新注册用户默认权限" json:"isDefault"` // 是否为新注册用户默认权限
	Users     []User     `gorm:"many2many:ga_user_roles" json:"-"`                               // 该角色下的用户（多对多）
	Children  []Role     `gorm:"-" json:"children"`
	Routes    []Route    `gorm:"-" json:"routes"`                      // 该角色的 api 权限
	Menus     []Menu     `gorm:"many2many:ga_role_menus" json:"menus"` // 该角色可使用的菜单
}

func (r Role) TableName() string {
	return "ga_roles"
}
