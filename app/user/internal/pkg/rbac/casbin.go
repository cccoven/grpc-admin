package rbac

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func NewCasbinEnforcer(db *gorm.DB, logger *zap.SugaredLogger) *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(db)
		// casbin RBAC 模型
		text := `
			[request_definition]
			r = sub, obj, act
			
			[policy_definition]
			p = sub, obj, act
			
			[role_definition]
			g = _, _
			
			[policy_effect]
			e = some(where (p.eft == allow))
			
			[matchers]
			m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
		`

		m, err := model.NewModelFromString(text)
		if err != nil {
			logger.Errorf("字符串载入 casbin 模型失败：%s", err.Error())
		}

		syncedEnforcer, _ = casbin.NewSyncedEnforcer(m, a)
	})

	_ = syncedEnforcer.LoadPolicy()

	return syncedEnforcer
}
