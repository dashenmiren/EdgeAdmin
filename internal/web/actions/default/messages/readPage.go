package messages

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type ReadPageAction struct {
	actionutils.ParentAction
}

func (this *ReadPageAction) RunPost(params struct {
	MessageIds []int64
}) {
	// 创建日志
	defer this.CreateLogInfo(codes.Message_LogReadMessages)

	_, err := this.RPC().MessageRPC().UpdateMessagesRead(this.AdminContext(), &pb.UpdateMessagesReadRequest{
		MessageIds: params.MessageIds,
		IsRead:     true,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
