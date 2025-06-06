// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .

package security

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeAdmin/internal/web/helpers"
)

type DismissXFFPromptAction struct {
	actionutils.ParentAction
}

func (this *DismissXFFPromptAction) RunPost(params struct{}) {
	helpers.DisableXFFPrompt()

	this.Success()
}
