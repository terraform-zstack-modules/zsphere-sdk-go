// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryNetworkServiceL3NetworkRef queries NetworkServiceL3NetworkRef list
func (cli *ZSClient) QueryNetworkServiceL3NetworkRef(ctx context.Context, params *param.QueryParam) ([]view.NetworkServiceL3NetworkRefInventoryView, error) {
	var resp []view.NetworkServiceL3NetworkRefInventoryView
	return resp, cli.List(ctx, "v1/l3-networks/network-services/refs", params, &resp)
}

// PageNetworkServiceL3NetworkRef Pagination
func (cli *ZSClient) PageNetworkServiceL3NetworkRef(ctx context.Context, params *param.QueryParam) ([]view.NetworkServiceL3NetworkRefInventoryView, int, error) {
	var networkServiceL3NetworkRefs []view.NetworkServiceL3NetworkRefInventoryView
	total, err := cli.Page(ctx, "v1/l3-networks/network-services/refs", params, &networkServiceL3NetworkRefs)
	return networkServiceL3NetworkRefs, total, err
}
