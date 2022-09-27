package conf

var AppConf = new(Conf)

type Conf struct {
	// Database 数据库配置
	Database struct {
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

	// Service rpc 服务配置
	Service struct {
		Name string
		Host string
		Port string
	}

	// Etcd 配置
	Etcd struct {
		Hosts []string
	}

	// Cache 缓存中间件配置
	Cache struct {
		Redis struct {
			Host     string
			Password string
			DB       int
		}
	}

	// Logger 日志组件配置
	Logger struct {
		Zap struct {
			OutputDir string
			Format    string
			Level     string
		}
	}
}
