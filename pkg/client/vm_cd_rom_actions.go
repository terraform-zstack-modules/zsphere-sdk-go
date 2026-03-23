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

// CreateVmCdRom creates VmCdRom
func (cli *ZSClient) CreateVmCdRom(ctx context.Context, params param.CreateVmCdRomParam) (*view.VmCdRomInventoryView, error) {
	resp := view.VmCdRomInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/cdroms"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVmCdRom updates VmCdRom
func (cli *ZSClient) UpdateVmCdRom(ctx context.Context, uuid string, params param.UpdateVmCdRomParam) (*view.VmCdRomInventoryView, error) {
	resp := view.VmCdRomInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances/cdroms", uuid, "actions", "inventory", map[string]interface{}{
		"updateVmCdRom": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteVmCdRom deletes VmCdRom
func (cli *ZSClient) DeleteVmCdRom(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances/cdroms", uuid, string(deleteMode))
}
// QueryVmCdRom queries VmCdRom list
func (cli *ZSClient) QueryVmCdRom(ctx context.Context, params *param.QueryParam) ([]view.VmCdRomInventoryView, error) {
	var resp []view.VmCdRomInventoryView
	return resp, cli.List(ctx, "v1/vm-instances/cdroms", params, &resp)
}

func (cli *ZSClient) GetVmCdRom(ctx context.Context, uuid string) (*view.VmCdRomInventoryView, error) {
	var resp view.VmCdRomInventoryView
	if err := cli.Get(ctx, "v1/vm-instances/cdroms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmCdRom Pagination
func (cli *ZSClient) PageVmCdRom(ctx context.Context, params *param.QueryParam) ([]view.VmCdRomInventoryView, int, error) {
	var vmCdRoms []view.VmCdRomInventoryView
	total, err := cli.Page(ctx, "v1/vm-instances/cdroms", params, &vmCdRoms)
	return vmCdRoms, total, err
}
