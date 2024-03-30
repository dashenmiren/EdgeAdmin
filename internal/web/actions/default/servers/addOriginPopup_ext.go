//go:build !plus

package servers

import (
	"github.com/dashenmiren/EdgeCommon/pkg/serverconfigs/ossconfigs"
	"github.com/iwind/TeaGo/maps"
)

func (this *AddOriginPopupAction) getOSSHook() {
	this.Data["ossTypes"] = []maps.Map{}
	this.Data["ossBucketParams"] = []maps.Map{}
	this.Data["ossForm"] = ""
}

func (this *AddOriginPopupAction) postOSSHook(protocol string) (config *ossconfigs.OSSConfig, goNext bool, err error) {
	goNext = true
	return
}
