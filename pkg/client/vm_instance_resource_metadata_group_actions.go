// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVmInstanceResourceMetadataGroup queries VmInstanceResourceMetadataGroup list
func (cli *ZSClient) QueryVmInstanceResourceMetadataGroup(ctx context.Context, params *param.QueryParam) ([]view.VmInstanceResourceMetadataGroupInventoryView, error) {
	var resp []view.VmInstanceResourceMetadataGroupInventoryView
	return resp, cli.List(ctx, "v1/vmInstance/resource/metadata/group", params, &resp)
}

func (cli *ZSClient) GetVmInstanceResourceMetadataGroup(ctx context.Context, uuid string) (*view.VmInstanceResourceMetadataGroupInventoryView, error) {
	var resp view.VmInstanceResourceMetadataGroupInventoryView
	if err := cli.Get(ctx, "v1/vmInstance/resource/metadata/group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmInstanceResourceMetadataGroup Pagination
func (cli *ZSClient) PageVmInstanceResourceMetadataGroup(ctx context.Context, params *param.QueryParam) ([]view.VmInstanceResourceMetadataGroupInventoryView, int, error) {
	var vmInstanceResourceMetadataGroups []view.VmInstanceResourceMetadataGroupInventoryView
	total, err := cli.Page(ctx, "v1/vmInstance/resource/metadata/group", params, &vmInstanceResourceMetadataGroups)
	return vmInstanceResourceMetadataGroups, total, err
}
