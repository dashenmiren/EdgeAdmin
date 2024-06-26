package ipbox

import (
	"strings"
	"time"

	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type AddIPAction struct {
	actionutils.ParentAction
}

func (this *AddIPAction) RunPost(params struct {
	ListId    int64
	Ip        string
	ExpiredAt int64
}) {
	var itemId int64 = 0

	defer func() {
		this.CreateLogInfo(codes.IPItem_LogCreateIPItem, params.ListId, itemId)
	}()

	var ipType = "ipv4"
	if strings.Contains(params.Ip, ":") {
		ipType = "ipv6"
	}

	if params.ExpiredAt <= 0 {
		params.ExpiredAt = time.Now().Unix() + 86400
	}

	createResp, err := this.RPC().IPItemRPC().CreateIPItem(this.AdminContext(), &pb.CreateIPItemRequest{
		IpListId:   params.ListId,
		IpFrom:     params.Ip,
		IpTo:       "",
		ExpiredAt:  params.ExpiredAt,
		Reason:     "从IPBox中加入名单",
		Type:       ipType,
		EventLevel: "critical",
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	itemId = createResp.IpItemId

	this.Success()
}
