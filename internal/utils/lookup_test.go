package utils_test

import (
	"testing"

	"github.com/dashenmiren/EdgeAdmin/internal/utils"
)

func TestLookupCNAME(t *testing.T) {
	for _, domain := range []string{"www.yun4s.cn", "example.com"} {
		result, err := utils.LookupCNAME(domain)
		t.Log(domain, "=>", result, err)
	}
}
