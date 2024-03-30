package headers

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/serverconfigs"
)

type OptionsAction struct {
	actionutils.ParentAction
}

func (this *OptionsAction) RunPost(params struct {
	Type string
}) {
	if params.Type == "request" {
		this.Data["headers"] = serverconfigs.AllHTTPCommonRequestHeaders
	} else if params.Type == "response" {
		this.Data["headers"] = serverconfigs.AllHTTPCommonResponseHeaders
	} else {
		this.Data["headers"] = []string{}
	}

	this.Success()
}
