package model

import "grpc-admin/common/pkg"

type UserSignInLog struct {
	pkg.GormModel
	UserID uint   `gorm:"column:user_id" json:"userID"`
	Ip     string `gorm:"column:ip" json:"ip"`
	Agent  string `json:"agent"` // 用户设备
}

func (u UserSignInLog) TableName() string {
	return "ga_user_signin_logs"
}
