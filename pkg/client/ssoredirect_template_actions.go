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

// CreateSSORedirectTemplate creates SSORedirectTemplate
func (cli *ZSClient) CreateSSORedirectTemplate(ctx context.Context, params param.CreateSSORedirectTemplateParam) (*view.SSORedirectTemplateInventoryView, error) {
	resp := view.SSORedirectTemplateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/create/sso/redirect/template/"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSSORedirectTemplate updates SSORedirectTemplate
func (cli *ZSClient) UpdateSSORedirectTemplate(ctx context.Context, params param.UpdateSSORedirectTemplateParam) (*view.SSORedirectTemplateInventoryView, error) {
	resp := view.SSORedirectTemplateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/update/sso/redirectTemplate"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteSSORedirectTemplate deletes SSORedirectTemplate
func (cli *ZSClient) DeleteSSORedirectTemplate(ctx context.Context, params param.DeleteSSORedirectTemplateParam) (*view.SSORedirectTemplateInventoryView, error) {
	resp := view.SSORedirectTemplateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/delete/sso/redirect/template"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
