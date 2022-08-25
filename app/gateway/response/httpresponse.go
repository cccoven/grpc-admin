package response

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"net/http"
)

const (
	// DefaultErrCode 默认网关错误状态码
	DefaultErrCode = http.StatusInternalServerError

	DefaultErrMsg = "请求失败"
)

type HttpResponse struct {
	Status  int    `json:"status"`
	Result  any    `json:"result"`
	Message string `json:"message"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, HttpResponse{
		Status:  http.StatusOK,
		Result:  data,
		Message: http.StatusText(http.StatusOK),
	})
}

func Error(c *gin.Context, status int, msg string) {
	c.JSON(http.StatusOK, HttpResponse{
		Status:  status,
		Result:  nil,
		Message: msg,
	})
}

func ParameterError(c *gin.Context, msg string) {
	Error(c, http.StatusBadRequest, msg)
}

func DefaultError(c *gin.Context) {
	Error(c, DefaultErrCode, DefaultErrMsg)
}

func RpcError(c *gin.Context, err error) {
	grpcStatus := status.Convert(err)
	code := int(grpcStatus.Code())
	msg := grpcStatus.Message()

	Error(c, code, msg)
}
