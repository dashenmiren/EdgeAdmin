// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .

package tasks

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
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
