//go:build !plus

package cluster

func (this *CreateNodeAction) findNodesQuota() (maxNodes int32, leftNodes int32, err error) {
	return
}
