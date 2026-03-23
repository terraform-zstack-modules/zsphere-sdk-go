// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QuerySharedBlockGroupPrimaryStorageHostRef queries SharedBlockGroupPrimaryStorageHostRef list
func (cli *ZSClient) QuerySharedBlockGroupPrimaryStorageHostRef(ctx context.Context, params *param.QueryParam) ([]view.SharedBlockGroupPrimaryStorageHostRefInventoryView, error) {
	var resp []view.SharedBlockGroupPrimaryStorageHostRefInventoryView
	return resp, cli.List(ctx, "v1/sharedblock-group/host-refs", params, &resp)
}

func (cli *ZSClient) GetSharedBlockGroupPrimaryStorageHostRef(ctx context.Context, uuid string) (*view.SharedBlockGroupPrimaryStorageHostRefInventoryView, error) {
	var resp view.SharedBlockGroupPrimaryStorageHostRefInventoryView
	if err := cli.Get(ctx, "v1/sharedblock-group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSharedBlockGroupPrimaryStorageHostRef Pagination
func (cli *ZSClient) PageSharedBlockGroupPrimaryStorageHostRef(ctx context.Context, params *param.QueryParam) ([]view.SharedBlockGroupPrimaryStorageHostRefInventoryView, int, error) {
	var sharedBlockGroupPrimaryStorageHostRefs []view.SharedBlockGroupPrimaryStorageHostRefInventoryView
	total, err := cli.Page(ctx, "v1/sharedblock-group/host-refs", params, &sharedBlockGroupPrimaryStorageHostRefs)
	return sharedBlockGroupPrimaryStorageHostRefs, total, err
}
