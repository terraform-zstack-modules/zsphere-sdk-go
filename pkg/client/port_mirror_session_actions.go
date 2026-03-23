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

// CreatePortMirrorSession creates PortMirrorSession
func (cli *ZSClient) CreatePortMirrorSession(ctx context.Context, params param.CreatePortMirrorSessionParam) (*view.PortMirrorSessionInventoryView, error) {
	resp := view.PortMirrorSessionInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/port-mirrors/sessions"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePortMirrorSession deletes PortMirrorSession
func (cli *ZSClient) DeletePortMirrorSession(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/port-mirrors/sessons", uuid, string(deleteMode))
}
// QueryPortMirrorSession queries PortMirrorSession list
func (cli *ZSClient) QueryPortMirrorSession(ctx context.Context, params *param.QueryParam) ([]view.PortMirrorSessionInventoryView, error) {
	var resp []view.PortMirrorSessionInventoryView
	return resp, cli.List(ctx, "v1/port-mirrors/sessions", params, &resp)
}

func (cli *ZSClient) GetPortMirrorSession(ctx context.Context, uuid string) (*view.PortMirrorSessionInventoryView, error) {
	var resp view.PortMirrorSessionInventoryView
	if err := cli.Get(ctx, "v1/port-mirrors/sessions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePortMirrorSession Pagination
func (cli *ZSClient) PagePortMirrorSession(ctx context.Context, params *param.QueryParam) ([]view.PortMirrorSessionInventoryView, int, error) {
	var portMirrorSessions []view.PortMirrorSessionInventoryView
	total, err := cli.Page(ctx, "v1/port-mirrors/sessions", params, &portMirrorSessions)
	return portMirrorSessions, total, err
}
