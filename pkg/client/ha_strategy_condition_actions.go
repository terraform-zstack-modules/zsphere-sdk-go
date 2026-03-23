// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateHaStrategyCondition updates HaStrategyCondition
func (cli *ZSClient) UpdateHaStrategyCondition(ctx context.Context, uuid string, params param.UpdateHaStrategyConditionParam) (*view.HaStrategyConditionInventoryView, error) {
	resp := view.HaStrategyConditionInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/ha-strategy-condition", uuid, "actions", "inventory", map[string]interface{}{
		"updateHaStrategyCondition": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
