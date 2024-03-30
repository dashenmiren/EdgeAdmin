package utils_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeAdmin/internal/utils"
)

func TestLookupCNAME(t *testing.T) {
	for _, domain := range []string{"www.yun4s.cn", "example.com"} {
		result, err := utils.LookupCNAME(domain)
		t.Log(domain, "=>", result, err)
	}
}
