package servers

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

// DeleteServersAction 删除一组网站
type DeleteServersAction struct {
	actionutils.ParentAction
}

func (this *DeleteServersAction) RunPost(params struct {
	ServerIds []int64
}) {
	defer this.CreateLogInfo(codes.Server_LogDeleteServers)

	_, err := this.RPC().ServerRPC().DeleteServers(this.AdminContext(), &pb.DeleteServersRequest{ServerIds: params.ServerIds})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
