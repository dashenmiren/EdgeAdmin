//go:build !plus

package cluster

func (this *CreateBatchAction) findNodesQuota() (maxNodes int32, leftNodes int32, err error) {
	return
}
