//go:build !plus

package ui

import "github.com/dashenmiren/EdgeCommon/pkg/systemconfigs"

func (this *IndexAction) filterConfig(config *systemconfigs.AdminUIConfig) {
	this.Data["supportModuleCDN"] = true
	this.Data["supportModuleNS"] = true
	this.Data["nsIsVisible"] = false
}
