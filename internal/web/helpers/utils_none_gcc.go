//go:build !gcc

package helpers

import (
	"net/http"
)

// filter request
func safeFilterRequest(req *http.Request) bool {
	return true
}
