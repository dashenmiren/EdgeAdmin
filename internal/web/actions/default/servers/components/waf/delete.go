package waf

import (	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	FirewallPolicyId int64
}) {
	// 日志
	defer this.CreateLogInfo(codes.WAFPolicy_LogDeleteWAFPolicy, params.FirewallPolicyId)

	countResp, err := this.RPC().NodeClusterRPC().CountAllEnabledNodeClustersWithHTTPFirewallPolicyId(this.AdminContext(), &pb.CountAllEnabledNodeClustersWithHTTPFirewallPolicyIdRequest{HttpFirewallPolicyId: params.FirewallPolicyId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if countResp.Count > 0 {
		this.Fail("此WAF策略正在被有些集群引用，请修改后再删除。")
	}

	_, err = this.RPC().HTTPFirewallPolicyRPC().DeleteHTTPFirewallPolicy(this.AdminContext(), &pb.DeleteHTTPFirewallPolicyRequest{HttpFirewallPolicyId: params.FirewallPolicyId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
