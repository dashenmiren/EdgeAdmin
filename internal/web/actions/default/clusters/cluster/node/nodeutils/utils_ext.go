//go:build !plus

package nodeutils

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/maps"
)

func filterMenuItems(menuItems []maps.Map, menuItem string, prefix string, query string, info *pb.FindEnabledNodeConfigInfoResponse, langCode string) []maps.Map {
	return menuItems
}
