package iplists

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/actions"
)

type UpdateAction struct {
	actionutils.ParentAction
}

func (this *UpdateAction) Init() {
	this.Nav("", "", "update")
}

func (this *UpdateAction) RunGet(params struct {
	ListId int64
}) {
	err := InitIPList(this.Parent(), params.ListId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Show()
}

func (this *UpdateAction) RunPost(params struct {
	ListId      int64
	Name        string
	Type        string
	Description string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo(codes.IPList_LogUpdateIPList, params.ListId)

	params.Must.
		Field("name", params.Name).
		Require("请输入名称")

	_, err := this.RPC().IPListRPC().UpdateIPList(this.AdminContext(), &pb.UpdateIPListRequest{
		IpListId:    params.ListId,
		Name:        params.Name,
		Code:        "",
		TimeoutJSON: nil,
		Description: params.Description,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
