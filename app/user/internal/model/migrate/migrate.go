package migrate

import (
	"gorm.io/gorm"
	"grpc-admin/app/user/internal/model"
	"log"
)

func Do(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.UserSignInLog{},
		&model.Role{},
		&model.UserRole{},
		&model.Menu{},
		&model.Route{},
	)
	if err != nil {
		log.Fatal("Failed to migrate: ", err)
	}
}
