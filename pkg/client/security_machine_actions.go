// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// DeleteSecurityMachine deletes SecurityMachine
func (cli *ZSClient) DeleteSecurityMachine(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/security-machines", uuid, string(deleteMode))
}
// UpdateSecurityMachine updates SecurityMachine
func (cli *ZSClient) UpdateSecurityMachine(ctx context.Context, uuid string, params param.UpdateSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/security-machines", uuid, "actions", "", map[string]interface{}{
		"updateSecurityMachine": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySecurityMachine queries SecurityMachine list
func (cli *ZSClient) QuerySecurityMachine(ctx context.Context, params *param.QueryParam) ([]view.SecurityMachineInventoryView, error) {
	var resp []view.SecurityMachineInventoryView
	return resp, cli.List(ctx, "v1/security-machines", params, &resp)
}

func (cli *ZSClient) GetSecurityMachine(ctx context.Context, uuid string) (*view.SecurityMachineInventoryView, error) {
	var resp view.SecurityMachineInventoryView
	if err := cli.Get(ctx, "v1/security-machine", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSecurityMachine Pagination
func (cli *ZSClient) PageSecurityMachine(ctx context.Context, params *param.QueryParam) ([]view.SecurityMachineInventoryView, int, error) {
	var securityMachines []view.SecurityMachineInventoryView
	total, err := cli.Page(ctx, "v1/security-machines", params, &securityMachines)
	return securityMachines, total, err
}
