//go:build !plus

package dashboard

func (this *IndexAction) checkPlus() bool {
	return false
}
