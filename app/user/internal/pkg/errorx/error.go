package errorx

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// ErrDefault rpc 服务默认错误码，接在 code.Codes 后面
	ErrDefault = 500

	// ErrResourceNoExist 目标资源不存在
	ErrResourceNoExist = 4000

	// ErrResourceAlreadyExist 目标资源已存在
	ErrResourceAlreadyExist = 4001

	// ErrResourceBeingUsed 目标资源正在被使用
	ErrResourceBeingUsed = 4002

	// ErrResourceNoPermission 没有目标资源权限
	ErrResourceNoPermission = 4003
)

var errMessages = map[int]string{
	ErrDefault:              "请求失败",
	ErrResourceNoExist:      "目标资源不存在",
	ErrResourceAlreadyExist: "目标资源已存在",
	ErrResourceBeingUsed:    "目标资源正在被使用",
	ErrResourceNoPermission: "没有目标资源权限",
}

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (r *Error) Error() string {
	return r.Msg
}

func New(code int, msg string) error {
	return &Error{code, msg}
}

func NewFromCode(code int) error {
	return &Error{Code: code, Msg: errMessages[code]}
}

func Default() error {
	return &Error{Code: ErrDefault, Msg: errMessages[ErrDefault]}
}

func StatusError(err error) error {
	e := err.(*Error)
	return status.Error(codes.Code(e.Code), e.Error())
}
