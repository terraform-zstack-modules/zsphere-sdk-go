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

// QuerySlbGroup queries SlbGroup list
func (cli *ZSClient) QuerySlbGroup(ctx context.Context, params *param.QueryParam) ([]view.SlbGroupInventoryView, error) {
	var resp []view.SlbGroupInventoryView
	return resp, cli.List(ctx, "v1/load-balancers/slb/groups", params, &resp)
}

func (cli *ZSClient) GetSlbGroup(ctx context.Context, uuid string) (*view.SlbGroupInventoryView, error) {
	var resp view.SlbGroupInventoryView
	if err := cli.Get(ctx, "v1/load-balancers/slb/groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSlbGroup Pagination
func (cli *ZSClient) PageSlbGroup(ctx context.Context, params *param.QueryParam) ([]view.SlbGroupInventoryView, int, error) {
	var slbGroups []view.SlbGroupInventoryView
	total, err := cli.Page(ctx, "v1/load-balancers/slb/groups", params, &slbGroups)
	return slbGroups, total, err
}
// CreateSlbGroup creates SlbGroup
func (cli *ZSClient) CreateSlbGroup(ctx context.Context, params param.CreateSlbGroupParam) (*view.SlbGroupInventoryView, error) {
	resp := view.SlbGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/load-balancers/slb/groups"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteSlbGroup deletes SlbGroup
func (cli *ZSClient) DeleteSlbGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/load-balancers/slb/group", uuid, string(deleteMode))
}
// UpdateSlbGroup updates SlbGroup
func (cli *ZSClient) UpdateSlbGroup(ctx context.Context, uuid string, params param.UpdateSlbGroupParam) (*view.SlbGroupInventoryView, error) {
	resp := view.SlbGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/load-balancers/slb/group", uuid, "actions", "inventory", map[string]interface{}{
		"updateSlbGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
