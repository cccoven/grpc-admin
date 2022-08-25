package migrate

import (
	"gorm.io/gorm"
	"grpc-admin/app/system/internal/model"
)

func Do(db *gorm.DB) {
	db.AutoMigrate(
		&model.Route{},
		&model.RouteGroup{},
		&model.Menu{},
	)
}
