package conds

import (
	"github.com/dashenmiren/EdgeAdmin/internal/configloaders"
	"github.com/dashenmiren/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Prefix("/servers/server/settings/conds").
			GetPost("/addGroupPopup", new(AddGroupPopupAction)).
			GetPost("/addCondPopup", new(AddCondPopupAction)).
			EndAll()
	})
}
