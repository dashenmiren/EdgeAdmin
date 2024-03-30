package compression

import (
	"github.com/dashenmiren/EdgeAdmin/internal/configloaders"
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/locationutils"
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/default/servers/serverutils"
	"github.com/dashenmiren/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Helper(locationutils.NewLocationHelper()).
			Helper(serverutils.NewServerHelper()).
			Data("tinyMenuItem", "compression").
			Prefix("/servers/server/settings/locations/compression").
			GetPost("", new(IndexAction)).
			EndAll()
	})
}
