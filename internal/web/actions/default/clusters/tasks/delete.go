package tasks

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	TaskId int64
}) {
	defer this.CreateLogInfo(codes.NodeTask_LogDeleteNodeTask, params.TaskId)

	_, err := this.RPC().NodeTaskRPC().DeleteNodeTask(this.AdminContext(), &pb.DeleteNodeTaskRequest{NodeTaskId: params.TaskId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
