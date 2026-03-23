// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// GetVmDns gets VmDns by uuid
func (cli *ZSClient) GetVmDns(ctx context.Context, uuid string) (*view.VmDnsInventoryView, error) {
	var resp view.VmDnsInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// SetVmDns operates on VmDns
func (cli *ZSClient) SetVmDns(ctx context.Context, vmInstanceUuid string, params param.SetVmDnsParam) (*view.VmDnsInventoryView, error) {
	resp := view.VmDnsInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, "actions", "inventories", map[string]interface{}{
		"setVmDns": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
