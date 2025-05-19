// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .
//go:build !plus

package clusterutils

import (
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
)

func (this *ClusterHelper) filterMenuItems1(items []maps.Map, info *pb.FindEnabledNodeClusterConfigInfoResponse, clusterIdString string, selectedItem string, actionPtr actions.ActionWrapper) []maps.Map {
	return items
}

func (this *ClusterHelper) filterMenuItems2(items []maps.Map, info *pb.FindEnabledNodeClusterConfigInfoResponse, clusterIdString string, selectedItem string, actionPtr actions.ActionWrapper) []maps.Map {
	return items
}
