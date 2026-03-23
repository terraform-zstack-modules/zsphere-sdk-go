// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryScsiLun queries ScsiLun list
func (cli *ZSClient) QueryScsiLun(ctx context.Context, params *param.QueryParam) ([]view.ScsiLunInventoryView, error) {
	var resp []view.ScsiLunInventoryView
	return resp, cli.List(ctx, "v1/storage-devices/scsi-lun/luns", params, &resp)
}

func (cli *ZSClient) GetScsiLun(ctx context.Context, uuid string) (*view.ScsiLunInventoryView, error) {
	var resp view.ScsiLunInventoryView
	if err := cli.Get(ctx, "v1/storage-devices/scsi-lun/luns", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageScsiLun Pagination
func (cli *ZSClient) PageScsiLun(ctx context.Context, params *param.QueryParam) ([]view.ScsiLunInventoryView, int, error) {
	var scsiLuns []view.ScsiLunInventoryView
	total, err := cli.Page(ctx, "v1/storage-devices/scsi-lun/luns", params, &scsiLuns)
	return scsiLuns, total, err
}
// UpdateScsiLun updates ScsiLun
func (cli *ZSClient) UpdateScsiLun(ctx context.Context, uuid string, params param.UpdateScsiLunParam) (*view.ScsiLunInventoryView, error) {
	resp := view.ScsiLunInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/storage-devices/scsi-lun", uuid, "actions", "inventory", map[string]interface{}{
		"updateScsiLun": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
