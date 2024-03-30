package access

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
			Prefix("/servers/server/settings/access").
			GetPost("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)).
			GetPost("/updatePopup", new(UpdatePopupAction)).
			Post("/random", new(RandomAction)).
			EndAll()
	})
}
