// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package clusters

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/maps"
)

type NodeOptionsAction struct {
	actionutils.ParentAction
}

func (this *NodeOptionsAction) RunPost(params struct {
	ClusterId int64
}) {
	resp, err := this.RPC().NodeRPC().FindAllEnabledNodesWithNodeClusterId(this.AdminContext(), &pb.FindAllEnabledNodesWithNodeClusterIdRequest{NodeClusterId: params.ClusterId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	var nodeMaps = []maps.Map{}
	for _, node := range resp.Nodes {
		nodeMaps = append(nodeMaps, maps.Map{
			"id":   node.Id,
			"name": node.Name,
		})
	}
	this.Data["nodes"] = nodeMaps

	this.Success()
}
