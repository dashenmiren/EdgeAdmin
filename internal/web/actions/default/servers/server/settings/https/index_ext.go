//go:build !plus

package https

func (this *IndexAction) checkSupportsHTTP3(clusterId int64) (bool, error) {
	return false, nil
}
