package tasks

import (
	"github.com/dashenmiren/EdgeAdmin/internal/configloaders"
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/default/clusters/clusterutils"
	"github.com/dashenmiren/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeDNS)).
			Helper(clusterutils.NewClustersHelper()).
			Prefix("/dns/tasks").
			GetPost("/listPopup", new(ListPopupAction)).
			Post("/check", new(CheckAction)).
			Post("/delete", new(DeleteAction)).
			Post("/deleteAll", new(DeleteAllAction)).
			EndAll()
	})
}
