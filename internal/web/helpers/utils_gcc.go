//go:build gcc

package helpers

import (
	"net/http"

	"github.com/dashenmiren/EdgeAdmin/internal/waf/injectionutils"
)

// filter request
func safeFilterRequest(req *http.Request) bool {
	return !injectionutils.DetectXSS(req.RequestURI, false) && !injectionutils.DetectSQLInjection(req.RequestURI, false)
}
