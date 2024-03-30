//go:build !plus

package origins

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/ossconfigs"
	"github.com/iwind/TeaGo/maps"
)

func (this *UpdatePopupAction) getOSSHook() {
	this.Data["ossTypes"] = []maps.Map{}
	this.Data["ossBucketParams"] = []maps.Map{}
	this.Data["ossForm"] = ""
}

func (this *UpdatePopupAction) postOSSHook(protocol string) (config *ossconfigs.OSSConfig, goNext bool, err error) {
	goNext = true
	return
}
