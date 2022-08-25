package rpcerror

import "google.golang.org/grpc/codes"

const (
	ErrDefault = 888
)

var errMessages = map[codes.Code]string{
	ErrDefault: "Request failed",
}

func New(code codes.Code) (codes.Code, string) {
	return code, errMessages[code]
}
