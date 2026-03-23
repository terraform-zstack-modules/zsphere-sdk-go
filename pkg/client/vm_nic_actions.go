// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"fmt"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// CreateVmNic creates VmNic
func (cli *ZSClient) CreateVmNic(ctx context.Context, params param.CreateVmNicParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/nics"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVmNic queries VmNic list
func (cli *ZSClient) QueryVmNic(ctx context.Context, params *param.QueryParam) ([]view.VmNicInventoryView, error) {
	var resp []view.VmNicInventoryView
	return resp, cli.List(ctx, "v1/vm-instances/nics", params, &resp)
}

func (cli *ZSClient) GetVmNic(ctx context.Context, uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.Get(ctx, "v1/vm-instances/nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmNic Pagination
func (cli *ZSClient) PageVmNic(ctx context.Context, params *param.QueryParam) ([]view.VmNicInventoryView, int, error) {
	var vmNics []view.VmNicInventoryView
	total, err := cli.Page(ctx, "v1/vm-instances/nics", params, &vmNics)
	return vmNics, total, err
}
// DeleteVmNic deletes VmNic
func (cli *ZSClient) DeleteVmNic(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/nics", uuid, string(deleteMode))
}
