package log

import (
	"github.com/dashenmiren/EdgeAdmin/internal/configloaders"
	"github.com/dashenmiren/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Helper(NewHelper()).
			Prefix("/servers/components/log").
			Get("", new(IndexAction)).
			EndAll()
	})
}
