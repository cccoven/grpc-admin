package model

import (
	"gorm.io/gorm"
	"time"
)

// RouteGroup 路由组
type RouteGroup struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `json:"name"`
}

func (r RouteGroup) TableName() string {
	return "ga_route_groups"
}
