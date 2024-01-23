package server

import (
	"context"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/log"
	"github.com/pkg/errors"
	runtime "k8s.io/cri-api/pkg/apis/runtime/v1"
	"path"
	"strings"
	"time"
)

func (c *criService) RestoreContainer(ctx context.Context, r *runtime.RestoreContainerRequest) (res *runtime.RestoreContainerResponse, err error) {
	//pathSlice := strings.Split(r.GetLocation(), "/")
	//containerName := pathSlice[len(pathSlice)-2]
	//unTarClonePodIdContainerNamePath := path.Join(r.GetClonePodId(), containerName)
	unTarClonePodIdContainerNamePath, _ := path.Split(strings.Split(r.GetLocation(), "/migration/")[1])
	if err := c.startContainer(ctx, r.GetContainerId(), unTarClonePodIdContainerNamePath, containerd.WithRestoreImagePath(r.GetLocation())); err != nil {
		return nil, errors.Wrap(err, "failed to restore container")
	}
	log.G(ctx).Infof("Restore Time :%v.", time.Now().UnixMilli())
	return &runtime.RestoreContainerResponse{}, nil
}
