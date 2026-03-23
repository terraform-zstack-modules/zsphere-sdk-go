// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVpcHaGroupNetworkServiceRef queries VpcHaGroupNetworkServiceRef list
func (cli *ZSClient) QueryVpcHaGroupNetworkServiceRef(ctx context.Context, params *param.QueryParam) ([]view.VpcHaGroupNetworkServiceRefInventoryView, error) {
	var resp []view.VpcHaGroupNetworkServiceRefInventoryView
	return resp, cli.List(ctx, "v1/vpc/hagroups/networkserviceref/", params, &resp)
}

func (cli *ZSClient) GetVpcHaGroupNetworkServiceRef(ctx context.Context, uuid string) (*view.VpcHaGroupNetworkServiceRefInventoryView, error) {
	var resp view.VpcHaGroupNetworkServiceRefInventoryView
	if err := cli.Get(ctx, "v1/vpc/hagroups/networkserviceref", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcHaGroupNetworkServiceRef Pagination
func (cli *ZSClient) PageVpcHaGroupNetworkServiceRef(ctx context.Context, params *param.QueryParam) ([]view.VpcHaGroupNetworkServiceRefInventoryView, int, error) {
	var vpcHaGroupNetworkServiceRefs []view.VpcHaGroupNetworkServiceRefInventoryView
	total, err := cli.Page(ctx, "v1/vpc/hagroups/networkserviceref/", params, &vpcHaGroupNetworkServiceRefs)
	return vpcHaGroupNetworkServiceRefs, total, err
}
