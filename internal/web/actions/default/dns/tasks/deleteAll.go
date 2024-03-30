package tasks

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/langs/codes"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
)

type DeleteAllAction struct {
	actionutils.ParentAction
}

func (this *DeleteAllAction) RunPost(params struct{}) {
	defer this.CreateLogInfo(codes.DNSTask_LogDeleteAllDNSTasks)

	_, err := this.RPC().DNSTaskRPC().DeleteAllDNSTasks(this.AdminContext(), &pb.DeleteAllDNSTasksRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
