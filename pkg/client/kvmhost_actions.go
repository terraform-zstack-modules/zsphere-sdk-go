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

// AddKVMHost adds KVMHost
func (cli *ZSClient) AddKVMHost(ctx context.Context, params param.AddKVMHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/hosts/kvm"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddKVMHostAsync Async
func (cli *ZSClient) AddKVMHostAsync(ctx context.Context, params param.AddKVMHostParam) (string, error) {

	resource := "hosts/kvm"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
// UpdateKVMHost updates KVMHost
func (cli *ZSClient) UpdateKVMHost(ctx context.Context, uuid string, params param.UpdateKVMHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts/kvm", uuid, "actions", "inventory", map[string]interface{}{
		"updateKVMHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
