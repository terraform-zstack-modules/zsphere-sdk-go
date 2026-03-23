// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateAlarmLabel updates AlarmLabel
func (cli *ZSClient) UpdateAlarmLabel(ctx context.Context, uuid string, params param.UpdateAlarmLabelParam) (*view.AlarmLabelInventoryView, error) {
	resp := view.AlarmLabelInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zwatch/alarms/labels", uuid, "actions", "inventory", map[string]interface{}{
		"updateAlarmLabel": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
