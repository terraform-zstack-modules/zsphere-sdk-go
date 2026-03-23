// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryIscsiLun queries IscsiLun list
func (cli *ZSClient) QueryIscsiLun(ctx context.Context, params *param.QueryParam) ([]view.IscsiLunInventoryView, error) {
	var resp []view.IscsiLunInventoryView
	return resp, cli.List(ctx, "v1/storage-devices/iscsi/luns", params, &resp)
}

func (cli *ZSClient) GetIscsiLun(ctx context.Context, uuid string) (*view.IscsiLunInventoryView, error) {
	var resp view.IscsiLunInventoryView
	if err := cli.Get(ctx, "v1/storage-devices/iscsi/luns", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIscsiLun Pagination
func (cli *ZSClient) PageIscsiLun(ctx context.Context, params *param.QueryParam) ([]view.IscsiLunInventoryView, int, error) {
	var iscsiLuns []view.IscsiLunInventoryView
	total, err := cli.Page(ctx, "v1/storage-devices/iscsi/luns", params, &iscsiLuns)
	return iscsiLuns, total, err
}
