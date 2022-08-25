package helper

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"grpc-admin/app/gateway/conf"
	"time"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

type CustomClaims struct {
	UserID   uint32
	Username string
	RoleIDs  []uint32
}

type Claims struct {
	CustomClaims
	jwt.StandardClaims
	BufferTime int64
}

func NewClaims(customClaims CustomClaims) Claims {
	return Claims{
		CustomClaims: customClaims,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                         // 签名生效时间
			ExpiresAt: time.Now().Unix() + conf.AppConf.Jwt.ExpiresTime, // 签名过期时间
			Issuer:    conf.AppConf.Jwt.Issuer,                          // 签名发行者
		},
		// 缓冲时间，缓冲时间内会获得新的 token 刷新令牌，此时一个用户会存在两个有效令牌，但是前端只留一个，另一个会丢失
		BufferTime: conf.AppConf.Jwt.BufferTime,
	}
}

func GenJwt(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString([]byte(conf.AppConf.Jwt.SigningKey))
	return s, err
}

func ParseJwt(t string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(t, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.AppConf.Jwt.SigningKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
