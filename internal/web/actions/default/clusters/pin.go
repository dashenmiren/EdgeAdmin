package clusters

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/langs/codes"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
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
