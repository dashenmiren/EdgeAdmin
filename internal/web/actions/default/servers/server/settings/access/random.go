package access

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/rands"
)

type RandomAction struct {
	actionutils.ParentAction
}

func (this *RandomAction) RunPost(params struct{}) {
	this.Data["random"] = rands.HexString(32)

	this.Success()
}
