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

// UpdateLoadBalancer updates LoadBalancer
func (cli *ZSClient) UpdateLoadBalancer(ctx context.Context, uuid string, params param.UpdateLoadBalancerParam) (*view.LoadBalancerInventoryView, error) {
	resp := view.LoadBalancerInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/load-balancers", uuid, "actions", "inventory", map[string]interface{}{
		"updateLoadBalancer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteLoadBalancer deletes LoadBalancer
func (cli *ZSClient) DeleteLoadBalancer(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/load-balancers", uuid, string(deleteMode))
}
// QueryLoadBalancer queries LoadBalancer list
func (cli *ZSClient) QueryLoadBalancer(ctx context.Context, params *param.QueryParam) ([]view.LoadBalancerInventoryView, error) {
	var resp []view.LoadBalancerInventoryView
	return resp, cli.List(ctx, "v1/load-balancers", params, &resp)
}

func (cli *ZSClient) GetLoadBalancer(ctx context.Context, uuid string) (*view.LoadBalancerInventoryView, error) {
	var resp view.LoadBalancerInventoryView
	if err := cli.Get(ctx, "v1/load-balancers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLoadBalancer Pagination
func (cli *ZSClient) PageLoadBalancer(ctx context.Context, params *param.QueryParam) ([]view.LoadBalancerInventoryView, int, error) {
	var loadBalancers []view.LoadBalancerInventoryView
	total, err := cli.Page(ctx, "v1/load-balancers", params, &loadBalancers)
	return loadBalancers, total, err
}
// RefreshLoadBalancer operates on LoadBalancer
func (cli *ZSClient) RefreshLoadBalancer(ctx context.Context, uuid string, params param.RefreshLoadBalancerParam) (*view.LoadBalancerInventoryView, error) {
	resp := view.LoadBalancerInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/load-balancers", uuid, "actions", "inventory", map[string]interface{}{
		"refreshLoadBalancer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateLoadBalancer creates LoadBalancer
func (cli *ZSClient) CreateLoadBalancer(ctx context.Context, params param.CreateLoadBalancerParam) (*view.LoadBalancerInventoryView, error) {
	resp := view.LoadBalancerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/load-balancers"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
