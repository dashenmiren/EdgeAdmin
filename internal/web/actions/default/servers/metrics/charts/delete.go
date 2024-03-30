package charts

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	ChartId int64
}) {
	defer this.CreateLogInfo(codes.MetricChart_LogDeleteMetricChart, params.ChartId)

	_, err := this.RPC().MetricChartRPC().DeleteMetricChart(this.AdminContext(), &pb.DeleteMetricChartRequest{MetricChartId: params.ChartId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
