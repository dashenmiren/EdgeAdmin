package charts

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/default/servers/metrics/metricutils"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
	"github.com/dashenmiren/EdgeCommon/pkg/serverconfigs"
	"github.com/iwind/TeaGo/maps"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "chart")
}

func (this *IndexAction) RunGet(params struct {
	ItemId int64
}) {
	_, err := metricutils.InitItem(this.Parent(), params.ItemId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	countResp, err := this.RPC().MetricChartRPC().CountEnabledMetricCharts(this.AdminContext(), &pb.CountEnabledMetricChartsRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	var count = countResp.Count
	page := this.NewPage(count)
	this.Data["page"] = page.AsHTML()

	chartsResp, err := this.RPC().MetricChartRPC().ListEnabledMetricCharts(this.AdminContext(), &pb.ListEnabledMetricChartsRequest{
		MetricItemId: params.ItemId,
		Offset:       page.Offset,
		Size:         page.Size,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	var charts = chartsResp.MetricCharts
	var chartMaps = []maps.Map{}
	for _, chart := range charts {
		chartMaps = append(chartMaps, maps.Map{
			"id":       chart.Id,
			"name":     chart.Name,
			"type":     chart.Type,
			"typeName": serverconfigs.FindMetricChartTypeName(chart.Type),
			"isOn":     chart.IsOn,
			"widthDiv": chart.WidthDiv,
		})
	}
	this.Data["charts"] = chartMaps

	this.Show()
}
