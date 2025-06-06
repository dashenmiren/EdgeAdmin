// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .

package origins

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type UpdateIsOnAction struct {
	actionutils.ParentAction
}

func (this *UpdateIsOnAction) RunPost(params struct {
	OriginId int64
	IsOn     bool
}) {
	defer this.CreateLogInfo(codes.ServerOrigin_LogUpdateOriginIsOn, params.OriginId)

	_, err := this.RPC().OriginRPC().UpdateOriginIsOn(this.AdminContext(), &pb.UpdateOriginIsOnRequest{
		OriginId: params.OriginId,
		IsOn:     params.IsOn,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
