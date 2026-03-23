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

// DeleteFlowCollector deletes FlowCollector
func (cli *ZSClient) DeleteFlowCollector(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/flowmeters/collectors", uuid, string(deleteMode))
}
// CreateFlowCollector creates FlowCollector
func (cli *ZSClient) CreateFlowCollector(ctx context.Context, flowMeterUuid string, params param.CreateFlowCollectorParam) (*view.FlowCollectorInventoryView, error) {
	resp := view.FlowCollectorInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/flowmeters/%s/collectors", flowMeterUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryFlowCollector queries FlowCollector list
func (cli *ZSClient) QueryFlowCollector(ctx context.Context, params *param.QueryParam) ([]view.FlowCollectorInventoryView, error) {
	var resp []view.FlowCollectorInventoryView
	return resp, cli.List(ctx, "v1/flowmeters/collectors", params, &resp)
}

func (cli *ZSClient) GetFlowCollector(ctx context.Context, uuid string) (*view.FlowCollectorInventoryView, error) {
	var resp view.FlowCollectorInventoryView
	if err := cli.Get(ctx, "v1/flowmeters/collectors", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFlowCollector Pagination
func (cli *ZSClient) PageFlowCollector(ctx context.Context, params *param.QueryParam) ([]view.FlowCollectorInventoryView, int, error) {
	var flowCollectors []view.FlowCollectorInventoryView
	total, err := cli.Page(ctx, "v1/flowmeters/collectors", params, &flowCollectors)
	return flowCollectors, total, err
}
// UpdateFlowCollector updates FlowCollector
func (cli *ZSClient) UpdateFlowCollector(ctx context.Context, uuid string, params param.UpdateFlowCollectorParam) (*view.FlowCollectorInventoryView, error) {
	resp := view.FlowCollectorInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/flowmeters/collectors", uuid, "actions", "inventory", map[string]interface{}{
		"updateFlowCollector": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
