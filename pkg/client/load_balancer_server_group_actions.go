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

// DeleteLoadBalancerServerGroup deletes LoadBalancerServerGroup
func (cli *ZSClient) DeleteLoadBalancerServerGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/load-balancers/servergroups", uuid, string(deleteMode))
}
// CreateLoadBalancerServerGroup creates LoadBalancerServerGroup
func (cli *ZSClient) CreateLoadBalancerServerGroup(ctx context.Context, loadBalancerUuid string, params param.CreateLoadBalancerServerGroupParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	resp := view.LoadBalancerServerGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/load-balancers/%s/servergroups", loadBalancerUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateLoadBalancerServerGroup updates LoadBalancerServerGroup
func (cli *ZSClient) UpdateLoadBalancerServerGroup(ctx context.Context, uuid string, params param.UpdateLoadBalancerServerGroupParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	resp := view.LoadBalancerServerGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/load-balancers/servergroups", uuid, "actions", "inventory", map[string]interface{}{
		"updateLoadBalancerServerGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryLoadBalancerServerGroup queries LoadBalancerServerGroup list
func (cli *ZSClient) QueryLoadBalancerServerGroup(ctx context.Context, params *param.QueryParam) ([]view.LoadBalancerServerGroupInventoryView, error) {
	var resp []view.LoadBalancerServerGroupInventoryView
	return resp, cli.List(ctx, "v1/load-balancers/servergroups", params, &resp)
}

func (cli *ZSClient) GetLoadBalancerServerGroup(ctx context.Context, uuid string) (*view.LoadBalancerServerGroupInventoryView, error) {
	var resp view.LoadBalancerServerGroupInventoryView
	if err := cli.Get(ctx, "v1/load-balancers/servergroups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLoadBalancerServerGroup Pagination
func (cli *ZSClient) PageLoadBalancerServerGroup(ctx context.Context, params *param.QueryParam) ([]view.LoadBalancerServerGroupInventoryView, int, error) {
	var loadBalancerServerGroups []view.LoadBalancerServerGroupInventoryView
	total, err := cli.Page(ctx, "v1/load-balancers/servergroups", params, &loadBalancerServerGroups)
	return loadBalancerServerGroups, total, err
}
