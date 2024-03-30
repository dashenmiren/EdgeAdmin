package apinodeutils_test

import (
	"runtime"
	"testing"

	"github.com/TeaOSLab/EdgeAdmin/internal/utils/apinodeutils"
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
