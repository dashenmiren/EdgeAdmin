package ssh

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type TestAction struct {
	actionutils.ParentAction
}

func (this *TestAction) RunPost(params struct {
	GrantId int64
	Host    string
	Port    int32
}) {
	resp, err := this.RPC().NodeGrantRPC().TestNodeGrant(this.AdminContext(), &pb.TestNodeGrantRequest{
		NodeGrantId: params.GrantId,
		Host:        params.Host,
		Port:        params.Port,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["isOk"] = resp.IsOk
	this.Data["error"] = resp.Error
	this.Success()
}
