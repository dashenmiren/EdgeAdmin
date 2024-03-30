package setup

import "github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"

var currentStatusText = ""

type StatusAction struct {
	actionutils.ParentAction
}

func (this *StatusAction) RunPost(params struct{}) {
	this.Data["statusText"] = currentStatusText
	this.Success()
}
