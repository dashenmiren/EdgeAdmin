package dns

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/rands"
)

type RandomNameAction struct {
	actionutils.ParentAction
}

func (this *RandomNameAction) RunPost(params struct{}) {
	this.Data["name"] = "cluster" + rands.HexString(8)

	this.Success()
}
