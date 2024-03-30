package logs

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeAdmin/internal/web/helpers"
	"github.com/dashenmiren/EdgeCommon/pkg/nodeconfigs"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type ReadLogsAction struct {
	actionutils.ParentAction
}

func (this *ReadLogsAction) RunPost(params struct {
	LogIds []int64

	NodeId int64
}) {
	_, err := this.RPC().NodeLogRPC().UpdateNodeLogsRead(this.AdminContext(), &pb.UpdateNodeLogsReadRequest{
		NodeLogIds: params.LogIds,
		NodeId:     params.NodeId,
		Role:       nodeconfigs.NodeRoleNode,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 通知左侧数字Badge更新
	helpers.NotifyNodeLogsCountChange()

	this.Success()
}
