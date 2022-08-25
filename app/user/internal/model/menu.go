package model

import (
	"gorm.io/gorm"
	"time"
)

// Menu 菜单表
type Menu struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Path      string         `gorm:"comment:路由组件路径" json:"path"`
	Name      string         `gorm:"comment:路由名称" json:"name"`
	Component string         `gorm:"comment:前端路由组件位置" json:"component"`
	Sort      int            `gorm:"排序标记" json:"sort"`
	ParentID  uint           `json:"parentID"`
	Children  []Menu         `gorm:"-" json:"children"`
}

func (m Menu) TableName() string {
	return "ga_menus"
}
