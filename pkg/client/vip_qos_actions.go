// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// DeleteVipQos deletes VipQos
func (cli *ZSClient) DeleteVipQos(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vips", uuid, string(deleteMode))
}
// GetVipQos gets VipQos by uuid
func (cli *ZSClient) GetVipQos(ctx context.Context, uuid string) (*view.VipQosInventoryView, error) {
	var resp view.VipQosInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vip", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// SetVipQos operates on VipQos
func (cli *ZSClient) SetVipQos(ctx context.Context, uuid string, params param.SetVipQosParam) (*view.VipQosInventoryView, error) {
	resp := view.VipQosInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vips", uuid, "actions", "inventory", map[string]interface{}{
		"setVipQos": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
