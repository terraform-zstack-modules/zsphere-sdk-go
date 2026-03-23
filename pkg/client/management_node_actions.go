// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryManagementNode queries ManagementNode list
func (cli *ZSClient) QueryManagementNode(ctx context.Context, params *param.QueryParam) ([]view.ManagementNodeInventoryView, error) {
	var resp []view.ManagementNodeInventoryView
	return resp, cli.List(ctx, "v1/management-nodes", params, &resp)
}

func (cli *ZSClient) GetManagementNode(ctx context.Context, uuid string) (*view.ManagementNodeInventoryView, error) {
	var resp view.ManagementNodeInventoryView
	if err := cli.Get(ctx, "v1/management-nodes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageManagementNode Pagination
func (cli *ZSClient) PageManagementNode(ctx context.Context, params *param.QueryParam) ([]view.ManagementNodeInventoryView, int, error) {
	var managementNodes []view.ManagementNodeInventoryView
	total, err := cli.Page(ctx, "v1/management-nodes", params, &managementNodes)
	return managementNodes, total, err
}
