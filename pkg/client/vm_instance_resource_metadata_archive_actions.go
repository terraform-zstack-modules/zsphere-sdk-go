// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVmInstanceResourceMetadataArchive queries VmInstanceResourceMetadataArchive list
func (cli *ZSClient) QueryVmInstanceResourceMetadataArchive(ctx context.Context, params *param.QueryParam) ([]view.VmInstanceResourceMetadataArchiveInventoryView, error) {
	var resp []view.VmInstanceResourceMetadataArchiveInventoryView
	return resp, cli.List(ctx, "v1/vmInstance/resource/metadata/archive", params, &resp)
}

func (cli *ZSClient) GetVmInstanceResourceMetadataArchive(ctx context.Context, uuid string) (*view.VmInstanceResourceMetadataArchiveInventoryView, error) {
	var resp view.VmInstanceResourceMetadataArchiveInventoryView
	if err := cli.Get(ctx, "v1/vmInstance/resource/metadata/archive", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmInstanceResourceMetadataArchive Pagination
func (cli *ZSClient) PageVmInstanceResourceMetadataArchive(ctx context.Context, params *param.QueryParam) ([]view.VmInstanceResourceMetadataArchiveInventoryView, int, error) {
	var vmInstanceResourceMetadataArchives []view.VmInstanceResourceMetadataArchiveInventoryView
	total, err := cli.Page(ctx, "v1/vmInstance/resource/metadata/archive", params, &vmInstanceResourceMetadataArchives)
	return vmInstanceResourceMetadataArchives, total, err
}
