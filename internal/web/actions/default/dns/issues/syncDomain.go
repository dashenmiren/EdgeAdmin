package issues

import (	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type SyncDomainAction struct {
	actionutils.ParentAction
}

func (this *SyncDomainAction) RunPost(params struct {
	DomainId int64
}) {
	// 记录日志
	defer this.CreateLogInfo(codes.DNS_LogSyncDomain, params.DomainId)

	// 执行同步
	resp, err := this.RPC().DNSDomainRPC().SyncDNSDomainData(this.AdminContext(), &pb.SyncDNSDomainDataRequest{DnsDomainId: params.DomainId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if resp.IsOk {
		this.Success()
	} else {
		this.Data["shouldFix"] = resp.ShouldFix
		this.Fail(resp.Error)
	}

	this.Success()
}
