// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package clusters

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type PinAction struct {
	actionutils.ParentAction
}

func (this *PinAction) RunPost(params struct {
	ClusterId int64
	IsPinned  bool
}) {
	if params.IsPinned {
		defer this.CreateLogInfo(codes.NodeCluster_LogPinCluster, params.ClusterId)
	} else {
		defer this.CreateLogInfo(codes.NodeCluster_LogUnpinCluster, params.ClusterId)
	}

	_, err := this.RPC().NodeClusterRPC().UpdateNodeClusterPinned(this.AdminContext(), &pb.UpdateNodeClusterPinnedRequest{
		NodeClusterId: params.ClusterId,
		IsPinned:      params.IsPinned,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
