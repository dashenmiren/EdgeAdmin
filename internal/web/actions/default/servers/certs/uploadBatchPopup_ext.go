//go:build !plus

package certs

func (this *UploadBatchPopupAction) maxFiles() int {
	return 20
}
