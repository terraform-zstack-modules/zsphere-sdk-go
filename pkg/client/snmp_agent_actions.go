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

// CreateSnmpAgent creates SnmpAgent
func (cli *ZSClient) CreateSnmpAgent(ctx context.Context, params param.CreateSnmpAgentParam) (*view.SnmpAgentInventoryView, error) {
	resp := view.SnmpAgentInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/snmp/agent"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StartSnmpAgent starts SnmpAgent
func (cli *ZSClient) StartSnmpAgent(ctx context.Context, params param.StartSnmpAgentParam) (*view.SnmpAgentInventoryView, error) {
	resp := view.SnmpAgentInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/snmp/agent/actions", "", "inventory", map[string]interface{}{
		"startSnmpAgent": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StopSnmpAgent stops SnmpAgent
func (cli *ZSClient) StopSnmpAgent(ctx context.Context, params param.StopSnmpAgentParam) (*view.SnmpAgentInventoryView, error) {
	resp := view.SnmpAgentInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/snmp/agent/actions", "", "inventory", map[string]interface{}{
		"stopSnmpAgent": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSnmpAgent updates SnmpAgent
func (cli *ZSClient) UpdateSnmpAgent(ctx context.Context, params param.UpdateSnmpAgentParam) (*view.SnmpAgentInventoryView, error) {
	resp := view.SnmpAgentInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/snmp/agent/actions", "", "inventory", map[string]interface{}{
		"updateSnmpAgent": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySnmpAgent queries SnmpAgent list
func (cli *ZSClient) QuerySnmpAgent(ctx context.Context, params *param.QueryParam) ([]view.SnmpAgentInventoryView, error) {
	var resp []view.SnmpAgentInventoryView
	return resp, cli.List(ctx, "v1/snmp/agent", params, &resp)
}

func (cli *ZSClient) GetSnmpAgent(ctx context.Context, uuid string) (*view.SnmpAgentInventoryView, error) {
	var resp view.SnmpAgentInventoryView
	if err := cli.Get(ctx, "v1/snmp/agent", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSnmpAgent Pagination
func (cli *ZSClient) PageSnmpAgent(ctx context.Context, params *param.QueryParam) ([]view.SnmpAgentInventoryView, int, error) {
	var snmpAgents []view.SnmpAgentInventoryView
	total, err := cli.Page(ctx, "v1/snmp/agent", params, &snmpAgents)
	return snmpAgents, total, err
}
