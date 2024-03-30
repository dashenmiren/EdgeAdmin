package groups

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type SortAction struct {
	actionutils.ParentAction
}

func (this *SortAction) RunPost(params struct {
	GroupIds []int64
}) {
	_, err := this.RPC().NodeGroupRPC().UpdateNodeGroupOrders(this.AdminContext(), &pb.UpdateNodeGroupOrdersRequest{NodeGroupIds: params.GroupIds})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 创建日志
	defer this.CreateLogInfo(codes.NodeGroup_LogSortNodeGroups)

	this.Success()
}
