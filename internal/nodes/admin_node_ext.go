//go:build !plus

package nodes

import (
	"github.com/dashenmiren/EdgeCommon/pkg/iplibrary"
	"github.com/iwind/TeaGo/logs"
)

// 启动IP库
func (this *AdminNode) startIPLibrary() {
	logs.Println("[NODE]initializing ip library ...")
	err := iplibrary.InitDefault()
	if err != nil {
		logs.Println("[NODE]initialize ip library failed: " + err.Error())
	}
}
