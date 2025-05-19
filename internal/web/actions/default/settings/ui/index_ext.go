// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .
//go:build !plus

package ui

import "github.com/dashenmiren/EdgeCommon/pkg/systemconfigs"

func (this *IndexAction) filterConfig(config *systemconfigs.AdminUIConfig) {
	this.Data["supportModuleCDN"] = true
	this.Data["supportModuleNS"] = true
	this.Data["nsIsVisible"] = false
}
