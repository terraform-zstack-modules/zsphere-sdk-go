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

// DeletePortGroup deletes PortGroup
func (cli *ZSClient) DeletePortGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l3-networks/port-group", uuid, string(deleteMode))
}
// CreatePortGroup creates PortGroup
func (cli *ZSClient) CreatePortGroup(ctx context.Context, params param.CreatePortGroupParam) (*view.PortGroupInventoryView, error) {
	resp := view.PortGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l3-networks/port-group"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdatePortGroup updates PortGroup
func (cli *ZSClient) UpdatePortGroup(ctx context.Context, uuid string, params param.UpdatePortGroupParam) (*view.PortGroupInventoryView, error) {
	resp := view.PortGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/l3-networks/port-group", uuid, "actions", "inventory", map[string]interface{}{
		"updatePortGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryPortGroup queries PortGroup list
func (cli *ZSClient) QueryPortGroup(ctx context.Context, params *param.QueryParam) ([]view.PortGroupInventoryView, error) {
	var resp []view.PortGroupInventoryView
	return resp, cli.List(ctx, "v1/l3-networks/port-group", params, &resp)
}

func (cli *ZSClient) GetPortGroup(ctx context.Context, uuid string) (*view.PortGroupInventoryView, error) {
	var resp view.PortGroupInventoryView
	if err := cli.Get(ctx, "v1/l3-networks/port-group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePortGroup Pagination
func (cli *ZSClient) PagePortGroup(ctx context.Context, params *param.QueryParam) ([]view.PortGroupInventoryView, int, error) {
	var portGroups []view.PortGroupInventoryView
	total, err := cli.Page(ctx, "v1/l3-networks/port-group", params, &portGroups)
	return portGroups, total, err
}
