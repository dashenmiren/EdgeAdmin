// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .

package apinodeutils_test

import (
	"runtime"
	"testing"

	"github.com/dashenmiren/EdgeAdmin/internal/utils/apinodeutils"
	_ "github.com/iwind/TeaGo/bootstrap"
)

func TestUpgrader_CanUpgrade(t *testing.T) {
	t.Log(apinodeutils.CanUpgrade("0.6.3", runtime.GOOS, runtime.GOARCH))
}

func TestUpgrader_Upgrade(t *testing.T) {
	var upgrader = apinodeutils.NewUpgrader(1)
	err := upgrader.Upgrade()
	if err != nil {
		t.Fatal(err)
	}
}
