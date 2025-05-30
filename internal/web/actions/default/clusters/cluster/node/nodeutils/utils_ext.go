// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .
//go:build !plus

package nodeutils

import (
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/maps"
)

func filterMenuItems(menuItems []maps.Map, menuItem string, prefix string, query string, info *pb.FindEnabledNodeConfigInfoResponse, langCode string) []maps.Map {
	return menuItems
}
