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

// DeletePolicyRouteTableRouteEntry deletes PolicyRouteTableRouteEntry
func (cli *ZSClient) DeletePolicyRouteTableRouteEntry(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/policy-routes/routes", uuid, string(deleteMode))
}
// QueryPolicyRouteTableRouteEntry queries PolicyRouteTableRouteEntry list
func (cli *ZSClient) QueryPolicyRouteTableRouteEntry(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteTableRouteEntryInventoryView, error) {
	var resp []view.PolicyRouteTableRouteEntryInventoryView
	return resp, cli.List(ctx, "v1/policy-routes/routes", params, &resp)
}

func (cli *ZSClient) GetPolicyRouteTableRouteEntry(ctx context.Context, uuid string) (*view.PolicyRouteTableRouteEntryInventoryView, error) {
	var resp view.PolicyRouteTableRouteEntryInventoryView
	if err := cli.Get(ctx, "v1/policy-routes/routes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePolicyRouteTableRouteEntry Pagination
func (cli *ZSClient) PagePolicyRouteTableRouteEntry(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteTableRouteEntryInventoryView, int, error) {
	var policyRouteTableRouteEntries []view.PolicyRouteTableRouteEntryInventoryView
	total, err := cli.Page(ctx, "v1/policy-routes/routes", params, &policyRouteTableRouteEntries)
	return policyRouteTableRouteEntries, total, err
}
// CreatePolicyRouteTableRouteEntry creates PolicyRouteTableRouteEntry
func (cli *ZSClient) CreatePolicyRouteTableRouteEntry(ctx context.Context, params param.CreatePolicyRouteTableRouteEntryParam) (*view.PolicyRouteTableRouteEntryInventoryView, error) {
	resp := view.PolicyRouteTableRouteEntryInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/policy-routes/routes"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
