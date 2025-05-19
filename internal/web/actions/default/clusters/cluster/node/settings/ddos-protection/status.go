// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package ddosProtection

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/default/nodes/nodeutils"
	"github.com/dashenmiren/EdgeCommon/pkg/messageconfigs"
)

type StatusAction struct {
	actionutils.ParentAction
}

func (this *StatusAction) RunPost(params struct {
	NodeId int64
}) {
	results, err := nodeutils.SendMessageToNodeIds(this.AdminContext(), []int64{params.NodeId}, messageconfigs.MessageCodeCheckLocalFirewall, &messageconfigs.CheckLocalFirewallMessage{
		Name: "nftables",
	}, 10)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["results"] = results
	this.Success()
}
