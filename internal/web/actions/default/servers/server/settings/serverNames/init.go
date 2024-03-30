package serverNames

import (
	"github.com/dashenmiren/EdgeAdmin/internal/configloaders"
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/default/servers/serverutils"
	"github.com/dashenmiren/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Helper(serverutils.NewServerHelper()).
			Data("mainTab", "setting").
			Data("secondMenuItem", "serverName").
			Prefix("/servers/server/settings/serverNames").
			GetPost("", new(IndexAction)).
			Post("/audit", new(AuditAction)).
			EndAll()
	})
}
