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

// QueryLoadBalancerListener queries LoadBalancerListener list
func (cli *ZSClient) QueryLoadBalancerListener(ctx context.Context, params *param.QueryParam) ([]view.LoadBalancerListenerInventoryView, error) {
	var resp []view.LoadBalancerListenerInventoryView
	return resp, cli.List(ctx, "v1/load-balancers/listeners", params, &resp)
}

func (cli *ZSClient) GetLoadBalancerListener(ctx context.Context, uuid string) (*view.LoadBalancerListenerInventoryView, error) {
	var resp view.LoadBalancerListenerInventoryView
	if err := cli.Get(ctx, "v1/load-balancers/listeners", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLoadBalancerListener Pagination
func (cli *ZSClient) PageLoadBalancerListener(ctx context.Context, params *param.QueryParam) ([]view.LoadBalancerListenerInventoryView, int, error) {
	var loadBalancerListeners []view.LoadBalancerListenerInventoryView
	total, err := cli.Page(ctx, "v1/load-balancers/listeners", params, &loadBalancerListeners)
	return loadBalancerListeners, total, err
}
// UpdateLoadBalancerListener updates LoadBalancerListener
func (cli *ZSClient) UpdateLoadBalancerListener(ctx context.Context, uuid string, params param.UpdateLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/load-balancers/listeners", uuid, "", "inventory", map[string]interface{}{
		"updateLoadBalancerListener": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ChangeLoadBalancerListener changes LoadBalancerListener
func (cli *ZSClient) ChangeLoadBalancerListener(ctx context.Context, uuid string, params param.ChangeLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/load-balancers/listeners", uuid, "actions", "inventory", map[string]interface{}{
		"changeLoadBalancerListener": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateLoadBalancerListener creates LoadBalancerListener
func (cli *ZSClient) CreateLoadBalancerListener(ctx context.Context, loadBalancerUuid string, params param.CreateLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/load-balancers/%s/listeners", loadBalancerUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteLoadBalancerListener deletes LoadBalancerListener
func (cli *ZSClient) DeleteLoadBalancerListener(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/load-balancers/listeners", uuid, string(deleteMode))
}
