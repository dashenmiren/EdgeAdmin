package users

import "github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"

type SelectPopupAction struct {
	actionutils.ParentAction
}

func (this *SelectPopupAction) Init() {
	this.Nav("", "", "")
}

func (this *SelectPopupAction) RunGet(params struct{}) {
	this.Show()
}
