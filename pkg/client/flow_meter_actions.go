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

// QueryFlowMeter queries FlowMeter list
func (cli *ZSClient) QueryFlowMeter(ctx context.Context, params *param.QueryParam) ([]view.FlowMeterInventoryView, error) {
	var resp []view.FlowMeterInventoryView
	return resp, cli.List(ctx, "v1/flowmeters", params, &resp)
}

func (cli *ZSClient) GetFlowMeter(ctx context.Context, uuid string) (*view.FlowMeterInventoryView, error) {
	var resp view.FlowMeterInventoryView
	if err := cli.Get(ctx, "v1/flowmeters", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFlowMeter Pagination
func (cli *ZSClient) PageFlowMeter(ctx context.Context, params *param.QueryParam) ([]view.FlowMeterInventoryView, int, error) {
	var flowMeters []view.FlowMeterInventoryView
	total, err := cli.Page(ctx, "v1/flowmeters", params, &flowMeters)
	return flowMeters, total, err
}
// CreateFlowMeter creates FlowMeter
func (cli *ZSClient) CreateFlowMeter(ctx context.Context, params param.CreateFlowMeterParam) (*view.FlowMeterInventoryView, error) {
	resp := view.FlowMeterInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/flowmeters"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteFlowMeter deletes FlowMeter
func (cli *ZSClient) DeleteFlowMeter(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/flowmeters", uuid, string(deleteMode))
}
// UpdateFlowMeter updates FlowMeter
func (cli *ZSClient) UpdateFlowMeter(ctx context.Context, uuid string, params param.UpdateFlowMeterParam) (*view.FlowMeterInventoryView, error) {
	resp := view.FlowMeterInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/flowmeters", uuid, "actions", "inventory", map[string]interface{}{
		"updateFlowMeter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
