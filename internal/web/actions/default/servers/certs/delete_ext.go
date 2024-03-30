//go:build !plus

package certs

func (this *DeleteAction) filterDelete(certId int64) error {
	return nil
}
