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

// DeleteZone deletes Zone
func (cli *ZSClient) DeleteZone(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zones", uuid, string(deleteMode))
}
// UpdateZone updates Zone
func (cli *ZSClient) UpdateZone(ctx context.Context, uuid string, params param.UpdateZoneParam) (*view.ZoneInventoryView, error) {
	resp := view.ZoneInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zones", uuid, "actions", "inventory", map[string]interface{}{
		"updateZone": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryZone queries Zone list
func (cli *ZSClient) QueryZone(ctx context.Context, params *param.QueryParam) ([]view.ZoneInventoryView, error) {
	var resp []view.ZoneInventoryView
	return resp, cli.List(ctx, "v1/zones", params, &resp)
}

func (cli *ZSClient) GetZone(ctx context.Context, uuid string) (*view.ZoneInventoryView, error) {
	var resp view.ZoneInventoryView
	if err := cli.Get(ctx, "v1/zones", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageZone Pagination
func (cli *ZSClient) PageZone(ctx context.Context, params *param.QueryParam) ([]view.ZoneInventoryView, int, error) {
	var zones []view.ZoneInventoryView
	total, err := cli.Page(ctx, "v1/zones", params, &zones)
	return zones, total, err
}
// CreateZone creates Zone
func (cli *ZSClient) CreateZone(ctx context.Context, params param.CreateZoneParam) (*view.ZoneInventoryView, error) {
	resp := view.ZoneInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zones"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
