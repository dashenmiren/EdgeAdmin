package server

import (
	teaconst "github.com/dashenmiren/EdgeAdmin/internal/const"
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
	"github.com/dashenmiren/EdgeCommon/pkg/serverconfigs"
	"strconv"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "index", "index")
	this.SecondMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	// TODO 等看板实现后，需要跳转到看板

	// TCP & UDP跳转到设置
	serverTypeResp, err := this.RPC().ServerRPC().FindEnabledServerType(this.AdminContext(), &pb.FindEnabledServerTypeRequest{ServerId: params.ServerId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	serverType := serverTypeResp.Type
	if serverType == serverconfigs.ServerTypeTCPProxy || serverType == serverconfigs.ServerTypeUDPProxy {
		this.RedirectURL("/servers/server/settings?serverId=" + strconv.FormatInt(params.ServerId, 10))
		return
	}

	// HTTP跳转到访问日志
	if teaconst.IsPlus {
		this.RedirectURL("/servers/server/boards?serverId=" + strconv.FormatInt(params.ServerId, 10))
	} else {
		this.RedirectURL("/servers/server/settings?serverId=" + strconv.FormatInt(params.ServerId, 10))
	}
}
