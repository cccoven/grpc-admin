package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"grpc-admin/app/gateway/conf"
	"grpc-admin/app/gateway/helper"
	"grpc-admin/app/gateway/pkg/cache"
	"grpc-admin/app/gateway/pkg/logger"
	"grpc-admin/app/gateway/request"
	"grpc-admin/app/gateway/response"
	"grpc-admin/app/gateway/rpc"
	"grpc-admin/app/thirdparty/thirdparty"
	"math/rand"
	"time"
)

type ThirdPartyApi struct {
	logger        *zap.SugaredLogger
	rds           *redis.Client
	thirdPartyRpc thirdparty.ThirdPartyClient
}

func NewThirdPartyApi() *ThirdPartyApi {
	return &ThirdPartyApi{
		logger:        logger.NewZapLogger(),
		rds:           cache.NewRedisCache(),
		thirdPartyRpc: thirdparty.NewThirdPartyClient(rpc.Discovery(conf.AppConf.ThirdPartyRpc.Name, conf.AppConf.ThirdPartyRpc.LoadBalanceMode)),
	}
}

// SendSMS 发送短信
func (t *ThirdPartyApi) SendSMS(c *gin.Context) {
	var req request.SendSMS
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}

	content := ""
	if req.Type == 1 { // 短信验证码
		randCode := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
		content = fmt.Sprintf("验证码：%s（10分钟内有效），用于账号登录。验证码设计账号安全，请勿告知他人。", randCode)

		// 将验证码存入缓存，10 分钟后过期
		key := fmt.Sprintf("smscode_%s", req.Phone)
		if err := t.rds.Set(context.Background(), key, randCode, 10*time.Minute).Err(); err != nil {
			t.logger.Error("存储验证码失败：", err.Error())
			response.DefaultError(c)
			return
		}
	} else {
		response.DefaultError(c)
	}

	// TODO 调用第三方 SDK 服务发送短信
	// _, err := t.thirdPartyRpc.SendSMS(context.Background(), &system.SendSMSRequest{Phone: req.Phone, Content: content})
	// if err != nil {
	// 	response.RpcError(c, err)
	// 	return
	// }
	fmt.Println(content)

	response.Success(c, content) // TODO 测试环境返回 content
}
