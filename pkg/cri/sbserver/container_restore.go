package sbserver

import (
	"context"
	runtime "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func (c *criService) RestoreContainer(ctx context.Context, r *runtime.RestoreContainerRequest) (res *runtime.RestoreContainerResponse, err error) {
	//TODO implement me
	panic("implement me")
}
