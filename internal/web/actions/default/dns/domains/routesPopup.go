package domains

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/maps"
)

type RoutesPopupAction struct {
	actionutils.ParentAction
}

func (this *RoutesPopupAction) Init() {
	this.Nav("", "", "")
}

func (this *RoutesPopupAction) RunGet(params struct {
	DomainId int64
}) {
	routesResp, err := this.RPC().DNSDomainRPC().FindAllDNSDomainRoutes(this.AdminContext(), &pb.FindAllDNSDomainRoutesRequest{DnsDomainId: params.DomainId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	routeMaps := []maps.Map{}
	for _, route := range routesResp.Routes {
		routeMaps = append(routeMaps, maps.Map{
			"name": route.Name,
			"code": route.Code,
		})
	}
	this.Data["routes"] = routeMaps

	this.Show()
}
