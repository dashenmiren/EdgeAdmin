package utils

import (
	"os/exec"
	"runtime"

	executils "github.com/dashenmiren/EdgeAdmin/internal/utils/exec"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/types"
)

func AddPortsToFirewall(ports []int) {
	for _, port := range ports {
		// Linux
		if runtime.GOOS == "linux" {
			// firewalld
			firewallCmd, _ := executils.LookPath("firewall-cmd")
			if len(firewallCmd) > 0 {
				err := exec.Command(firewallCmd, "--add-port="+types.String(port)+"/tcp").Run()
				if err == nil {
					logs.Println("ADMIN_NODE", "add port '"+types.String(port)+"' to firewalld")

					_ = exec.Command(firewallCmd, "--add-port="+types.String(port)+"/tcp", "--permanent").Run()
				}
			}
		}
	}
}
