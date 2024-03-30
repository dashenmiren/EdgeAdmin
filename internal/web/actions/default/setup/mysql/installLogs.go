package mysql

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/default/setup/mysql/mysqlinstallers/utils"
)

type InstallLogsAction struct {
	actionutils.ParentAction
}

func (this *InstallLogsAction) RunPost(params struct{}) {
	this.Data["logs"] = utils.SharedLogger.ReadAll()
	this.Success()
}
