package acme

import (
	"github.com/dashenmiren/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
	"github.com/iwind/TeaGo/maps"
)

type UserOptionsAction struct {
	actionutils.ParentAction
}

func (this *UserOptionsAction) RunPost(params struct {
	PlatformUserId int64
}) {
	// 获取所有可用的用户
	usersResp, err := this.RPC().ACMEUserRPC().FindAllACMEUsers(this.AdminContext(), &pb.FindAllACMEUsersRequest{
		AdminId: 0,
		UserId:  params.PlatformUserId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	var userMaps = []maps.Map{}
	for _, user := range usersResp.AcmeUsers {
		description := user.Description
		if len(description) > 0 {
			description = "（" + description + "）"
		}

		userMaps = append(userMaps, maps.Map{
			"id":           user.Id,
			"description":  description,
			"email":        user.Email,
			"providerCode": user.AcmeProviderCode,
		})
	}
	this.Data["users"] = userMaps

	this.Success()
}
