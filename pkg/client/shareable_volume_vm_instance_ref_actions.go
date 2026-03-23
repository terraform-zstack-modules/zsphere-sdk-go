// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryShareableVolumeVmInstanceRef queries ShareableVolumeVmInstanceRef list
func (cli *ZSClient) QueryShareableVolumeVmInstanceRef(ctx context.Context, params *param.QueryParam) ([]view.ShareableVolumeVmInstanceRefInventoryView, error) {
	var resp []view.ShareableVolumeVmInstanceRefInventoryView
	return resp, cli.List(ctx, "v1/volumes/vm-instances/refs", params, &resp)
}

// PageShareableVolumeVmInstanceRef Pagination
func (cli *ZSClient) PageShareableVolumeVmInstanceRef(ctx context.Context, params *param.QueryParam) ([]view.ShareableVolumeVmInstanceRefInventoryView, int, error) {
	var shareableVolumeVmInstanceRefs []view.ShareableVolumeVmInstanceRefInventoryView
	total, err := cli.Page(ctx, "v1/volumes/vm-instances/refs", params, &shareableVolumeVmInstanceRefs)
	return shareableVolumeVmInstanceRefs, total, err
}
