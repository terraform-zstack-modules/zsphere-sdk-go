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

// DeleteVpcHaGroup deletes VpcHaGroup
func (cli *ZSClient) DeleteVpcHaGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vpc/hagroups", uuid, string(deleteMode))
}
// CreateVpcHaGroup creates VpcHaGroup
func (cli *ZSClient) CreateVpcHaGroup(ctx context.Context, params param.CreateVpcHaGroupParam) (*view.VpcHaGroupInventoryView, error) {
	resp := view.VpcHaGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpc/hagroups"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVpcHaGroup updates VpcHaGroup
func (cli *ZSClient) UpdateVpcHaGroup(ctx context.Context, uuid string, params param.UpdateVpcHaGroupParam) (*view.VpcHaGroupInventoryView, error) {
	resp := view.VpcHaGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vpc/hagroups", uuid, "actions", "inventory", map[string]interface{}{
		"updateVpcHaGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVpcHaGroup queries VpcHaGroup list
func (cli *ZSClient) QueryVpcHaGroup(ctx context.Context, params *param.QueryParam) ([]view.VpcHaGroupInventoryView, error) {
	var resp []view.VpcHaGroupInventoryView
	return resp, cli.List(ctx, "v1/vpc/hagroups", params, &resp)
}

func (cli *ZSClient) GetVpcHaGroup(ctx context.Context, uuid string) (*view.VpcHaGroupInventoryView, error) {
	var resp view.VpcHaGroupInventoryView
	if err := cli.Get(ctx, "v1/vpc/hagroups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcHaGroup Pagination
func (cli *ZSClient) PageVpcHaGroup(ctx context.Context, params *param.QueryParam) ([]view.VpcHaGroupInventoryView, int, error) {
	var vpcHaGroups []view.VpcHaGroupInventoryView
	total, err := cli.Page(ctx, "v1/vpc/hagroups", params, &vpcHaGroups)
	return vpcHaGroups, total, err
}
