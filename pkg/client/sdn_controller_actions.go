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

// RemoveSdnController removes SdnController
func (cli *ZSClient) RemoveSdnController(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/sdn-controllers", uuid, string(deleteMode))
}
// AddSdnController adds SdnController
func (cli *ZSClient) AddSdnController(ctx context.Context, params param.AddSdnControllerParam) (*view.SdnControllerInventoryView, error) {
	resp := view.SdnControllerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sdn-controllers"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSdnController updates SdnController
func (cli *ZSClient) UpdateSdnController(ctx context.Context, uuid string, params param.UpdateSdnControllerParam) (*view.SdnControllerInventoryView, error) {
	resp := view.SdnControllerInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/sdn-controllers", uuid, "actions", "inventory", map[string]interface{}{
		"updateSdnController": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySdnController queries SdnController list
func (cli *ZSClient) QuerySdnController(ctx context.Context, params *param.QueryParam) ([]view.SdnControllerInventoryView, error) {
	var resp []view.SdnControllerInventoryView
	return resp, cli.List(ctx, "v1/sdn-controllers", params, &resp)
}

func (cli *ZSClient) GetSdnController(ctx context.Context, uuid string) (*view.SdnControllerInventoryView, error) {
	var resp view.SdnControllerInventoryView
	if err := cli.Get(ctx, "v1/sdn-controllers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSdnController Pagination
func (cli *ZSClient) PageSdnController(ctx context.Context, params *param.QueryParam) ([]view.SdnControllerInventoryView, int, error) {
	var sdnControllers []view.SdnControllerInventoryView
	total, err := cli.Page(ctx, "v1/sdn-controllers", params, &sdnControllers)
	return sdnControllers, total, err
}
