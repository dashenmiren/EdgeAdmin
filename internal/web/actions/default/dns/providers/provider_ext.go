//go:build !plus

package providers

import (
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/maps"
)

func (this *ProviderAction) readEdgeDNS(provider *pb.DNSProvider, apiParams maps.Map) (maps.Map, error) {
	return maps.Map{}, nil
}
