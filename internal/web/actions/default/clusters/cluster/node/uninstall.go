// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .

package node

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type UninstallAction struct {
	actionutils.ParentAction
}

func (this *UninstallAction) RunPost(params struct {
	NodeId int64
}) {
	resp, err := this.RPC().NodeRPC().UninstallNode(this.AdminContext(), &pb.UninstallNodeRequest{NodeId: params.NodeId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 创建日志
	defer this.CreateLogInfo(codes.Node_LogUninstallNodeRemotely, params.NodeId)

	if resp.IsOk {
		this.Success()
	}

	this.Fail("执行失败：" + resp.Error)
}
