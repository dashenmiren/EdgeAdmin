// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package remoteAddr

import (
	"encoding/json"
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/dao"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
	"github.com/dashenmiren/EdgeCommon/pkg/serverconfigs"
	"github.com/iwind/TeaGo/actions"
	"regexp"
	"strings"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {

}

func (this *IndexAction) RunGet(params struct {
	LocationId int64
}) {
	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithLocationId(this.AdminContext(), params.LocationId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["webId"] = webConfig.Id
	this.Data["remoteAddrConfig"] = webConfig.RemoteAddr

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	WebId          int64
	RemoteAddrJSON []byte

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	var remoteAddrConfig = &serverconfigs.HTTPRemoteAddrConfig{}
	err := json.Unmarshal(params.RemoteAddrJSON, remoteAddrConfig)
	if err != nil {
		this.Fail("参数校验失败：" + err.Error())
		return
	}

	remoteAddrConfig.Value = strings.TrimSpace(remoteAddrConfig.Value)

	switch remoteAddrConfig.Type {
	case serverconfigs.HTTPRemoteAddrTypeRequestHeader:
		if len(remoteAddrConfig.RequestHeaderName) == 0 {
			this.FailField("requestHeaderName", "请输入请求报头")
			return
		}
		if !regexp.MustCompile(`^[\w-_,]+$`).MatchString(remoteAddrConfig.RequestHeaderName) {
			this.FailField("requestHeaderName", "请求报头中只能含有数字、英文字母、下划线、中划线")
			return
		}
		remoteAddrConfig.Value = "${header." + remoteAddrConfig.RequestHeaderName + "}"
	case serverconfigs.HTTPRemoteAddrTypeVariable:
		if len(remoteAddrConfig.Value) == 0 {
			this.FailField("value", "请输入自定义变量")
			return
		}
	}

	err = remoteAddrConfig.Init()
	if err != nil {
		this.Fail("配置校验失败：" + err.Error())
		return
	}

	remoteAddrJSON, err := json.Marshal(remoteAddrConfig)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	_, err = this.RPC().HTTPWebRPC().UpdateHTTPWebRemoteAddr(this.AdminContext(), &pb.UpdateHTTPWebRemoteAddrRequest{
		HttpWebId:      params.WebId,
		RemoteAddrJSON: remoteAddrJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
