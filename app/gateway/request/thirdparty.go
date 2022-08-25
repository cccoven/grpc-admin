package request

type SendSMS struct {
	Phone string `json:"phone" validate:"required"`
	// 短信类型：1 登录验证码
	Type  int    `json:"type" validate:"required"`
}
