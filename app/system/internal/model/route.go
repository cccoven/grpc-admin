package model

import (
	"gorm.io/gorm"
	"time"
)

// Route 接口路由表
type Route struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Path      string         `json:"path"`
	Method    string         `json:"method"`
	Desc      string         `json:"desc"`
	GroupID   uint           `gorm:"column:group_id" json:"groupID"`
	Group     RouteGroup     `gorm:"foreignKey:GroupID;references:ID"`
}

func (r Route) TableName() string {
	return "ga_routes"
}
