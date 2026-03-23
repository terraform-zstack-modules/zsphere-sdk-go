// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateAlertDataAck updates AlertDataAck
func (cli *ZSClient) UpdateAlertDataAck(ctx context.Context, alertDataUuid string, params param.UpdateAlertDataAckParam) (*view.AlertDataAckInventoryView, error) {
	resp := view.AlertDataAckInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zwatch/alert-histories/acknowledgments", alertDataUuid, "actions", "inventory", map[string]interface{}{
		"updateAlertDataAck": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAlertDataAck queries AlertDataAck list
func (cli *ZSClient) QueryAlertDataAck(ctx context.Context, params *param.QueryParam) ([]view.AlertDataAckInventoryView, error) {
	var resp []view.AlertDataAckInventoryView
	return resp, cli.List(ctx, "v1/zwatch/alert-histories/acknowledgments", params, &resp)
}

func (cli *ZSClient) GetAlertDataAck(ctx context.Context, uuid string) (*view.AlertDataAckInventoryView, error) {
	var resp view.AlertDataAckInventoryView
	if err := cli.Get(ctx, "v1/zwatch/alert-histories/acknowledgments", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAlertDataAck Pagination
func (cli *ZSClient) PageAlertDataAck(ctx context.Context, params *param.QueryParam) ([]view.AlertDataAckInventoryView, int, error) {
	var alertDataAcks []view.AlertDataAckInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/alert-histories/acknowledgments", params, &alertDataAcks)
	return alertDataAcks, total, err
}
