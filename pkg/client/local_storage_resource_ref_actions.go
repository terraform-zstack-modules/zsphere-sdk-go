// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryLocalStorageResourceRef queries LocalStorageResourceRef list
func (cli *ZSClient) QueryLocalStorageResourceRef(ctx context.Context, params *param.QueryParam) ([]view.LocalStorageResourceRefInventoryView, error) {
	var resp []view.LocalStorageResourceRefInventoryView
	return resp, cli.List(ctx, "v1/primary-storage/local-storage/resource-refs", params, &resp)
}

// PageLocalStorageResourceRef Pagination
func (cli *ZSClient) PageLocalStorageResourceRef(ctx context.Context, params *param.QueryParam) ([]view.LocalStorageResourceRefInventoryView, int, error) {
	var localStorageResourceRefs []view.LocalStorageResourceRefInventoryView
	total, err := cli.Page(ctx, "v1/primary-storage/local-storage/resource-refs", params, &localStorageResourceRefs)
	return localStorageResourceRefs, total, err
}
