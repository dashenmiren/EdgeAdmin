package transfer

import (
	"github.com/dashenmiren/EdgeAdmin/internal/configloaders"
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/default/settings/settingutils"
	"github.com/dashenmiren/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeSetting)).
			Helper(settingutils.NewAdvancedHelper("transfer")).
			Prefix("/settings/transfer").
			Get("", new(IndexAction)).
			Post("/validateAPI", new(ValidateAPIAction)).
			Post("/updateHosts", new(UpdateHostsAction)).
			Post("/upgradeNodes", new(UpgradeNodesAction)).
			Post("/statNodes", new(StatNodesAction)).
			EndAll()
	})
}
