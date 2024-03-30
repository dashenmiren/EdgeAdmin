package logs

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeAdmin/internal/web/helpers"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type FixAllAction struct {
	actionutils.ParentAction
}

func (this *FixAllAction) RunPost(params struct {
}) {
	defer this.CreateLogInfo(codes.NodeLog_LogFixAllLogs)

	_, err := this.RPC().NodeLogRPC().FixAllNodeLogs(this.AdminContext(), &pb.FixAllNodeLogsRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 通知左侧数字Badge更新
	helpers.NotifyNodeLogsCountChange()

	this.Success()
}
