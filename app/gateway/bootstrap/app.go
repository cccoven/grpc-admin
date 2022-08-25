package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"grpc-admin/app/gateway/conf"
	"log"
)


func Run() {
	r := gin.Default()
	
	registerRouter(r)

	log.Fatal(r.Run(fmt.Sprintf(":%d", conf.AppConf.Server.Port)))
}
