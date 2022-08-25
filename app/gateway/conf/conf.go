package conf

var AppConf = new(Conf)

type Conf struct {
	Server struct {
		Port int
	}

	Jwt struct {
		SigningKey  string
		Issuer      string
		ExpiresTime int64
		BufferTime  int64
	}

	Etcd struct {
		Hosts []string
		Key   string
	}

	// Logger 日志组件配置
	Logger struct {
		Zap struct {
			OutputDir string
			Format    string
			Level     string
		}
	}

	Cache struct {
		Redis struct {
			Host     string
			Password string
			DB       int
		}
	}

	UserRpc struct {
		Name            string
		LoadBalanceMode string
	}

	ThirdPartyRpc struct {
		Name            string
		LoadBalanceMode string
	}

	SystemRpc struct {
		Name            string
		LoadBalanceMode string
	}
}
