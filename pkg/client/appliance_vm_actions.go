// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryApplianceVm queries ApplianceVm list
func (cli *ZSClient) QueryApplianceVm(ctx context.Context, params *param.QueryParam) ([]view.ApplianceVmInventoryView, error) {
	var resp []view.ApplianceVmInventoryView
	return resp, cli.List(ctx, "v1/vm-instances/appliances", params, &resp)
}

func (cli *ZSClient) GetApplianceVm(ctx context.Context, uuid string) (*view.ApplianceVmInventoryView, error) {
	var resp view.ApplianceVmInventoryView
	if err := cli.Get(ctx, "v1/vm-instances/appliances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageApplianceVm Pagination
func (cli *ZSClient) PageApplianceVm(ctx context.Context, params *param.QueryParam) ([]view.ApplianceVmInventoryView, int, error) {
	var applianceVms []view.ApplianceVmInventoryView
	total, err := cli.Page(ctx, "v1/vm-instances/appliances", params, &applianceVms)
	return applianceVms, total, err
}
