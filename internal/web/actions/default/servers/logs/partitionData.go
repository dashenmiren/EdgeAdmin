package logs

import (
	"strings"

	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

// PartitionDataAction 读取分区表
type PartitionDataAction struct {
	actionutils.ParentAction
}

func (this *PartitionDataAction) RunPost(params struct {
	Day string
}) {
	var day = params.Day
	day = strings.ReplaceAll(day, "-", "")

	resp, err := this.RPC().HTTPAccessLogRPC().FindHTTPAccessLogPartitions(this.AdminContext(), &pb.FindHTTPAccessLogPartitionsRequest{
		Day: day,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	if len(resp.Partitions) > 0 {
		this.Data["partitions"] = resp.Partitions
	} else {
		this.Data["partitions"] = []int32{}
	}

	this.Success()
}
