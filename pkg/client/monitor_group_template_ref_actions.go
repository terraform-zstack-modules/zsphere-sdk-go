// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryMonitorGroupTemplateRef queries MonitorGroupTemplateRef list
func (cli *ZSClient) QueryMonitorGroupTemplateRef(ctx context.Context, params *param.QueryParam) ([]view.MonitorGroupTemplateRefInventoryView, error) {
	var resp []view.MonitorGroupTemplateRefInventoryView
	return resp, cli.List(ctx, "v1/zwatch/monitorgroups/monitortemplates/refs", params, &resp)
}

// PageMonitorGroupTemplateRef Pagination
func (cli *ZSClient) PageMonitorGroupTemplateRef(ctx context.Context, params *param.QueryParam) ([]view.MonitorGroupTemplateRefInventoryView, int, error) {
	var monitorGroupTemplateRefs []view.MonitorGroupTemplateRefInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/monitorgroups/monitortemplates/refs", params, &monitorGroupTemplateRefs)
	return monitorGroupTemplateRefs, total, err
}
