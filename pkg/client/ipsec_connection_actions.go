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

// ReconnectIPsecConnection operates on IPsecConnection
func (cli *ZSClient) ReconnectIPsecConnection(ctx context.Context, uuid string, params param.ReconnectIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/ipsec", uuid, "actions", "inventory", map[string]interface{}{
		"reconnectIPsecConnection": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ChangeIPsecConnection changes IPsecConnection
func (cli *ZSClient) ChangeIPsecConnection(ctx context.Context, uuid string, params param.ChangeIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/ipsec/config", uuid, "", "inventory", map[string]interface{}{
		"changeIPsecConnection": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateIPsecConnection updates IPsecConnection
func (cli *ZSClient) UpdateIPsecConnection(ctx context.Context, uuid string, params param.UpdateIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/ipsec", uuid, "", "inventory", map[string]interface{}{
		"updateIPsecConnection": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteIPsecConnection deletes IPsecConnection
func (cli *ZSClient) DeleteIPsecConnection(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ipsec", uuid, string(deleteMode))
}
// CreateIPsecConnection creates IPsecConnection
func (cli *ZSClient) CreateIPsecConnection(ctx context.Context, params param.CreateIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/ipsec"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
