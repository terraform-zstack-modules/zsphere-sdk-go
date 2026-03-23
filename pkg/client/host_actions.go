// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateHost updates Host
func (cli *ZSClient) UpdateHost(ctx context.Context, uuid string, params param.UpdateHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts", uuid, "actions", "inventory", map[string]interface{}{
		"updateHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ReconnectHost operates on Host
func (cli *ZSClient) ReconnectHost(ctx context.Context, uuid string, params param.ReconnectHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts", uuid, "actions", "inventory", map[string]interface{}{
		"reconnectHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteHost deletes Host
func (cli *ZSClient) DeleteHost(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hosts", uuid, string(deleteMode))
}
// QueryHost queries Host list
func (cli *ZSClient) QueryHost(ctx context.Context, params *param.QueryParam) ([]view.HostInventoryView, error) {
	var resp []view.HostInventoryView
	return resp, cli.List(ctx, "v1/hosts", params, &resp)
}

func (cli *ZSClient) GetHost(ctx context.Context, uuid string) (*view.HostInventoryView, error) {
	var resp view.HostInventoryView
	if err := cli.Get(ctx, "v1/hosts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHost Pagination
func (cli *ZSClient) PageHost(ctx context.Context, params *param.QueryParam) ([]view.HostInventoryView, int, error) {
	var hosts []view.HostInventoryView
	total, err := cli.Page(ctx, "v1/hosts", params, &hosts)
	return hosts, total, err
}
