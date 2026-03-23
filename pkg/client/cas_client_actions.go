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

// CreateCasClient creates CasClient
func (cli *ZSClient) CreateCasClient(ctx context.Context, params param.CreateCasClientParam) (*view.CasClientInventoryView, error) {
	resp := view.CasClientInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/create/cas/client"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateCasClient updates CasClient
func (cli *ZSClient) UpdateCasClient(ctx context.Context, params param.UpdateCasClientParam) (*view.CasClientInventoryView, error) {
	resp := view.CasClientInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/update/cas/client"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
