package database

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type TruncateTableAction struct {
	actionutils.ParentAction
}

func (this *TruncateTableAction) RunPost(params struct {
	Table string
}) {
	defer this.CreateLogInfo(codes.Database_LogTruncateTable, params.Table)

	_, err := this.RPC().DBRPC().TruncateDBTable(this.AdminContext(), &pb.TruncateDBTableRequest{DbTable: params.Table})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
