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

// QueryResourceConfig queries ResourceConfig list
func (cli *ZSClient) QueryResourceConfig(ctx context.Context, params *param.QueryParam) ([]view.ResourceConfigInventoryView, error) {
	var resp []view.ResourceConfigInventoryView
	return resp, cli.List(ctx, "v1/resource-configurations", params, &resp)
}

// PageResourceConfig Pagination
func (cli *ZSClient) PageResourceConfig(ctx context.Context, params *param.QueryParam) ([]view.ResourceConfigInventoryView, int, error) {
	var resourceConfigs []view.ResourceConfigInventoryView
	total, err := cli.Page(ctx, "v1/resource-configurations", params, &resourceConfigs)
	return resourceConfigs, total, err
}
// UpdateResourceConfig updates ResourceConfig
func (cli *ZSClient) UpdateResourceConfig(ctx context.Context, category string, name string, resourceUuid string, params param.UpdateResourceConfigParam) (*view.ResourceConfigInventoryView, error) {
	resp := view.ResourceConfigInventoryView{}
	err := cli.PutWithSpec(ctx, "v1/resource-configurations", category, fmt.Sprintf("%s/%s/actions", name, resourceUuid), "inventory", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteResourceConfig deletes ResourceConfig
func (cli *ZSClient) DeleteResourceConfig(ctx context.Context, category string, name string, resourceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/resource-configurations", category, fmt.Sprintf("%s/%s", name, resourceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
