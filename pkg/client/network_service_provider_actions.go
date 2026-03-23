// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryNetworkServiceProvider queries NetworkServiceProvider list
func (cli *ZSClient) QueryNetworkServiceProvider(ctx context.Context, params *param.QueryParam) ([]view.NetworkServiceProviderInventoryView, error) {
	var resp []view.NetworkServiceProviderInventoryView
	return resp, cli.List(ctx, "v1/network-services/providers", params, &resp)
}

// PageNetworkServiceProvider Pagination
func (cli *ZSClient) PageNetworkServiceProvider(ctx context.Context, params *param.QueryParam) ([]view.NetworkServiceProviderInventoryView, int, error) {
	var networkServiceProviders []view.NetworkServiceProviderInventoryView
	total, err := cli.Page(ctx, "v1/network-services/providers", params, &networkServiceProviders)
	return networkServiceProviders, total, err
}
