package cache

import (
	"github.com/dashenmiren/EdgeAdmin/internal/rpc"
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
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
