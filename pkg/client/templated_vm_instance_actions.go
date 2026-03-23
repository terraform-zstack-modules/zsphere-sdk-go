// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateTemplatedVmInstance updates TemplatedVmInstance
func (cli *ZSClient) UpdateTemplatedVmInstance(ctx context.Context, uuid string, params param.UpdateTemplatedVmInstanceParam) (*view.TemplatedVmInstanceInventoryView, error) {
	resp := view.TemplatedVmInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances/templatedVmInstance", uuid, "actions", "inventory", map[string]interface{}{
		"updateTemplatedVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteTemplatedVmInstance deletes TemplatedVmInstance
func (cli *ZSClient) DeleteTemplatedVmInstance(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances/templatedVmInstance", uuid, string(deleteMode))
}
// QueryTemplatedVmInstance queries TemplatedVmInstance list
func (cli *ZSClient) QueryTemplatedVmInstance(ctx context.Context, params *param.QueryParam) ([]view.TemplatedVmInstanceInventoryView, error) {
	var resp []view.TemplatedVmInstanceInventoryView
	return resp, cli.List(ctx, "v1/vm-instances/templatedVmInstance", params, &resp)
}

func (cli *ZSClient) GetTemplatedVmInstance(ctx context.Context, uuid string) (*view.TemplatedVmInstanceInventoryView, error) {
	var resp view.TemplatedVmInstanceInventoryView
	if err := cli.Get(ctx, "v1/vm-instances/templatedVmInstance", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTemplatedVmInstance Pagination
func (cli *ZSClient) PageTemplatedVmInstance(ctx context.Context, params *param.QueryParam) ([]view.TemplatedVmInstanceInventoryView, int, error) {
	var templatedVmInstances []view.TemplatedVmInstanceInventoryView
	total, err := cli.Page(ctx, "v1/vm-instances/templatedVmInstance", params, &templatedVmInstances)
	return templatedVmInstances, total, err
}
