// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryGuestToolsState queries GuestToolsState list
func (cli *ZSClient) QueryGuestToolsState(ctx context.Context, params *param.QueryParam) ([]view.GuestToolsStateInventoryView, error) {
	var resp []view.GuestToolsStateInventoryView
	return resp, cli.List(ctx, "v1/guesttools", params, &resp)
}

func (cli *ZSClient) GetGuestToolsState(ctx context.Context, uuid string) (*view.GuestToolsStateInventoryView, error) {
	var resp view.GuestToolsStateInventoryView
	if err := cli.Get(ctx, "v1/guesttools/state", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageGuestToolsState Pagination
func (cli *ZSClient) PageGuestToolsState(ctx context.Context, params *param.QueryParam) ([]view.GuestToolsStateInventoryView, int, error) {
	var guestToolsStates []view.GuestToolsStateInventoryView
	total, err := cli.Page(ctx, "v1/guesttools", params, &guestToolsStates)
	return guestToolsStates, total, err
}
// UpdateGuestToolsState updates GuestToolsState
func (cli *ZSClient) UpdateGuestToolsState(ctx context.Context, vmInstanceUuid string, params param.UpdateGuestToolsStateParam) (*view.GuestToolsStateInventoryView, error) {
	resp := view.GuestToolsStateInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, "guesttools-state", "inventory", map[string]interface{}{
		"updateGuestToolsState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
