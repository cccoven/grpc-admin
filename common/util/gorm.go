package util

import "gorm.io/gorm"

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		offset := (page - 1) * pageSize
		return db.Limit(pageSize).Offset(offset)
	}
}
