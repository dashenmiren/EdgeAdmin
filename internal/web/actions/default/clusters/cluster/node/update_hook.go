//go:build !plus

package node

func (this *UpdateAction) CanUpdateLevel(level int32) bool {
	return level <= 1
}
