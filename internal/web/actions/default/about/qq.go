package about

import "github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"

type QqAction struct {
	actionutils.ParentAction
}

func (this *QqAction) Init() {
	this.Nav("", "", "")
}

func (this *QqAction) RunGet(params struct{}) {
	this.Show()
}
