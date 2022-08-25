package pkg

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type GormModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type GormConfig struct {
	Source       string
	User         string
	Password     string
	Host         string
	DbName       string
	Charset      string
	ParseTime    bool
	Loc          string
	MaxIdleConns int
	MaxOpenConns int
}

func NewGorm(c GormConfig) *gorm.DB {
	var db *gorm.DB
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
		c.User,
		c.Password,
		c.Host,
		c.DbName,
		c.Charset,
		c.ParseTime,
		c.Loc,
	)

	if c.Source == "mysql" {
		db = newMySqlGorm(dsn)
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	
	return db
}

func newMySqlGorm(dsn string) *gorm.DB {
	mysqlConf := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         188,   // string 类型默认大小
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConf))
	if err != nil{
		log.Fatalf("Failed to open database: %s", err)
	}
	
	return db
}
