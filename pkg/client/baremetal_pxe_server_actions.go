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

// ReconnectBaremetalPxeServer operates on BaremetalPxeServer
func (cli *ZSClient) ReconnectBaremetalPxeServer(ctx context.Context, uuid string, params param.ReconnectBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	resp := view.BaremetalPxeServerInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/pxeservers", uuid, "actions", "inventory", map[string]interface{}{
		"reconnectBaremetalPxeServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteBaremetalPxeServer deletes BaremetalPxeServer
func (cli *ZSClient) DeleteBaremetalPxeServer(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/baremetal/pxeservers", uuid, string(deleteMode))
}
// StopBaremetalPxeServer stops BaremetalPxeServer
func (cli *ZSClient) StopBaremetalPxeServer(ctx context.Context, uuid string, params param.StopBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	resp := view.BaremetalPxeServerInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/pxeservers", uuid, "actions", "inventory", map[string]interface{}{
		"stopBaremetalPxeServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateBaremetalPxeServer updates BaremetalPxeServer
func (cli *ZSClient) UpdateBaremetalPxeServer(ctx context.Context, uuid string, params param.UpdateBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	resp := view.BaremetalPxeServerInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/pxeservers", uuid, "actions", "inventory", map[string]interface{}{
		"updateBaremetalPxeServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StartBaremetalPxeServer starts BaremetalPxeServer
func (cli *ZSClient) StartBaremetalPxeServer(ctx context.Context, uuid string, params param.StartBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	resp := view.BaremetalPxeServerInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/pxeservers", uuid, "actions", "inventory", map[string]interface{}{
		"startBaremetalPxeServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryBaremetalPxeServer queries BaremetalPxeServer list
func (cli *ZSClient) QueryBaremetalPxeServer(ctx context.Context, params *param.QueryParam) ([]view.BaremetalPxeServerInventoryView, error) {
	var resp []view.BaremetalPxeServerInventoryView
	return resp, cli.List(ctx, "v1/baremetal/pxeservers", params, &resp)
}

func (cli *ZSClient) GetBaremetalPxeServer(ctx context.Context, uuid string) (*view.BaremetalPxeServerInventoryView, error) {
	var resp view.BaremetalPxeServerInventoryView
	if err := cli.Get(ctx, "v1/baremetal/pxeservers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBaremetalPxeServer Pagination
func (cli *ZSClient) PageBaremetalPxeServer(ctx context.Context, params *param.QueryParam) ([]view.BaremetalPxeServerInventoryView, int, error) {
	var baremetalPxeServers []view.BaremetalPxeServerInventoryView
	total, err := cli.Page(ctx, "v1/baremetal/pxeservers", params, &baremetalPxeServers)
	return baremetalPxeServers, total, err
}
// CreateBaremetalPxeServer creates BaremetalPxeServer
func (cli *ZSClient) CreateBaremetalPxeServer(ctx context.Context, params param.CreateBaremetalPxeServerParam) (*view.BaremetalPxeServerInventoryView, error) {
	resp := view.BaremetalPxeServerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/baremetal/pxeservers"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
