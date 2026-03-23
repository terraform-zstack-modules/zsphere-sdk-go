// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryEventLog queries EventLog list
func (cli *ZSClient) QueryEventLog(ctx context.Context, params *param.QueryParam) ([]view.EventLogInventoryView, error) {
	var resp []view.EventLogInventoryView
	return resp, cli.List(ctx, "v1/eventlogs", params, &resp)
}

func (cli *ZSClient) GetEventLog(ctx context.Context, uuid string) (*view.EventLogInventoryView, error) {
	var resp view.EventLogInventoryView
	if err := cli.Get(ctx, "v1/eventlogs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEventLog Pagination
func (cli *ZSClient) PageEventLog(ctx context.Context, params *param.QueryParam) ([]view.EventLogInventoryView, int, error) {
	var eventLogs []view.EventLogInventoryView
	total, err := cli.Page(ctx, "v1/eventlogs", params, &eventLogs)
	return eventLogs, total, err
}
