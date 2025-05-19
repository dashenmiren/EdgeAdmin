// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package ocsp

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/langs/codes"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
)

type ResetAction struct {
	actionutils.ParentAction
}

func (this *ResetAction) RunPost(params struct {
	CertIds []int64
}) {
	defer this.CreateLogInfo(codes.SSLCert_LogOCSPResetOCSPStatus)

	_, err := this.RPC().SSLCertRPC().ResetSSLCertsWithOCSPError(this.AdminContext(), &pb.ResetSSLCertsWithOCSPErrorRequest{SslCertIds: params.CertIds})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
