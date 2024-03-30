package logs

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Data("teaMenu", "servers").
			Data("teaSubMenu", "log").
			Prefix("/servers/logs").
			Get("", new(IndexAction)).
			GetPost("/settings", new(SettingsAction)).
			Post("/partitionData", new(PartitionDataAction)).
			Post("/hasLogs", new(HasLogsAction)).
			EndAll()
	})
}
