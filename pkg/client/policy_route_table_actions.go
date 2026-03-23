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

// CreatePolicyRouteTable creates PolicyRouteTable
func (cli *ZSClient) CreatePolicyRouteTable(ctx context.Context, params param.CreatePolicyRouteTableParam) (*view.PolicyRouteTableInventoryView, error) {
	resp := view.PolicyRouteTableInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/policy-routes/tables"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePolicyRouteTable deletes PolicyRouteTable
func (cli *ZSClient) DeletePolicyRouteTable(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/policy-routes/tables", uuid, string(deleteMode))
}
// QueryPolicyRouteTable queries PolicyRouteTable list
func (cli *ZSClient) QueryPolicyRouteTable(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteTableInventoryView, error) {
	var resp []view.PolicyRouteTableInventoryView
	return resp, cli.List(ctx, "v1/policy-routes/tables", params, &resp)
}

func (cli *ZSClient) GetPolicyRouteTable(ctx context.Context, uuid string) (*view.PolicyRouteTableInventoryView, error) {
	var resp view.PolicyRouteTableInventoryView
	if err := cli.Get(ctx, "v1/policy-routes/tables", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePolicyRouteTable Pagination
func (cli *ZSClient) PagePolicyRouteTable(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteTableInventoryView, int, error) {
	var policyRouteTables []view.PolicyRouteTableInventoryView
	total, err := cli.Page(ctx, "v1/policy-routes/tables", params, &policyRouteTables)
	return policyRouteTables, total, err
}
