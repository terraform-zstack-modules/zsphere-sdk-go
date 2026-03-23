// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryHostOsCategory queries HostOsCategory list
func (cli *ZSClient) QueryHostOsCategory(ctx context.Context, params *param.QueryParam) ([]view.HostOsCategoryInventoryView, error) {
	var resp []view.HostOsCategoryInventoryView
	return resp, cli.List(ctx, "v1/hosts/os/category", params, &resp)
}

// PageHostOsCategory Pagination
func (cli *ZSClient) PageHostOsCategory(ctx context.Context, params *param.QueryParam) ([]view.HostOsCategoryInventoryView, int, error) {
	var hostOsCategories []view.HostOsCategoryInventoryView
	total, err := cli.Page(ctx, "v1/hosts/os/category", params, &hostOsCategories)
	return hostOsCategories, total, err
}
