package waf

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type CountAction struct {
	actionutils.ParentAction
}

func (this *CountAction) RunPost(params struct{}) {
	countResp, err := this.RPC().HTTPFirewallPolicyRPC().CountAllEnabledHTTPFirewallPolicies(this.AdminContext(), &pb.CountAllEnabledHTTPFirewallPoliciesRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["count"] = countResp.Count

	this.Success()
}
