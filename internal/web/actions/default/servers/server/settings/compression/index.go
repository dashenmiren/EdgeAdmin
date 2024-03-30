package compression

import (
	"encoding/json"

	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/langs/codes"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/types"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("compression")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	// 服务分组设置
	groupResp, err := this.RPC().ServerGroupRPC().FindEnabledServerGroupConfigInfo(this.AdminContext(), &pb.FindEnabledServerGroupConfigInfoRequest{
		ServerId: params.ServerId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["hasGroupConfig"] = groupResp.HasCompressionConfig
	this.Data["groupSettingURL"] = "/servers/groups/group/settings/compression?groupId=" + types.String(groupResp.ServerGroupId)

	// WebId
	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithServerId(this.AdminContext(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["webId"] = webConfig.Id

	this.Data["compressionConfig"] = webConfig.Compression

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	WebId           int64
	CompressionJSON []byte

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo(codes.ServerCompression_LogUpdateCompressionSettings, params.WebId)

	// 校验配置
	var compressionConfig = &serverconfigs.HTTPCompressionConfig{}
	err := json.Unmarshal(params.CompressionJSON, compressionConfig)
	if err != nil {
		this.Fail("配置校验失败：" + err.Error())
	}

	err = compressionConfig.Init()
	if err != nil {
		this.Fail("配置校验失败：" + err.Error())
	}

	_, err = this.RPC().HTTPWebRPC().UpdateHTTPWebCompression(this.AdminContext(), &pb.UpdateHTTPWebCompressionRequest{
		HttpWebId:       params.WebId,
		CompressionJSON: params.CompressionJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
