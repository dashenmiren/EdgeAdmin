// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package metrics

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	ItemId int64
}) {
	defer this.CreateLogInfo(codes.MetricItem_LogDeleteMetricItem)

	_, err := this.RPC().MetricItemRPC().DeleteMetricItem(this.AdminContext(), &pb.DeleteMetricItemRequest{MetricItemId: params.ItemId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
