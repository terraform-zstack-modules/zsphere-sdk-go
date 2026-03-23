// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryCephOsdGroup queries CephOsdGroup list
func (cli *ZSClient) QueryCephOsdGroup(ctx context.Context, params *param.QueryParam) ([]view.CephOsdGroupInventoryView, error) {
	var resp []view.CephOsdGroupInventoryView
	return resp, cli.List(ctx, "v1/primary-storage/ceph/osdgroups", params, &resp)
}

func (cli *ZSClient) GetCephOsdGroup(ctx context.Context, uuid string) (*view.CephOsdGroupInventoryView, error) {
	var resp view.CephOsdGroupInventoryView
	if err := cli.Get(ctx, "v1/primary-storage/ceph/osdgroups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCephOsdGroup Pagination
func (cli *ZSClient) PageCephOsdGroup(ctx context.Context, params *param.QueryParam) ([]view.CephOsdGroupInventoryView, int, error) {
	var cephOsdGroups []view.CephOsdGroupInventoryView
	total, err := cli.Page(ctx, "v1/primary-storage/ceph/osdgroups", params, &cephOsdGroups)
	return cephOsdGroups, total, err
}
