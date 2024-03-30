package cache

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/rpc"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
)

func InitMenu(parent *actionutils.ParentAction) error {
	rpcClient, err := rpc.SharedRPC()
	if err != nil {
		return err
	}

	countTasksResp, err := rpcClient.HTTPCacheTaskRPC().CountDoingHTTPCacheTasks(parent.AdminContext(), &pb.CountDoingHTTPCacheTasksRequest{})
	if err != nil {
		return err
	}

	parent.Data["countDoingTasks"] = countTasksResp.Count
	return nil
}
