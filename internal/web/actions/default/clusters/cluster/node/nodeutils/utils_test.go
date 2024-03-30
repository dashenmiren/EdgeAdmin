package nodeutils_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/clusters/cluster/node/nodeutils"
	_ "github.com/iwind/TeaGo/bootstrap"
)

func TestInstallLocalNode(t *testing.T) {
	err := nodeutils.InstallLocalNode()
	if err != nil {
		t.Fatal(err)
	}
}
