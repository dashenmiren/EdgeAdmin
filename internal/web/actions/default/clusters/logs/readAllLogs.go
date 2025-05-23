// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package logs

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeAdmin/internal/web/helpers"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type ReadAllLogsAction struct {
	actionutils.ParentAction
}

func (this *ReadAllLogsAction) RunPost(params struct {
	LogIds []int64
}) {
	_, err := this.RPC().NodeLogRPC().UpdateAllNodeLogsRead(this.AdminContext(), &pb.UpdateAllNodeLogsReadRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 通知左侧数字Badge更新
	helpers.NotifyNodeLogsCountChange()

	this.Success()
}
