package messages

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type BadgeAction struct {
	actionutils.ParentAction
}

func (this *BadgeAction) RunPost(params struct{}) {
	countResp, err := this.RPC().MessageRPC().CountUnreadMessages(this.AdminContext(), &pb.CountUnreadMessagesRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["count"] = countResp.Count

	this.Success()
}
