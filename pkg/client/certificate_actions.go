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

// DeleteCertificate deletes Certificate
func (cli *ZSClient) DeleteCertificate(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/certificates", uuid, string(deleteMode))
}
// UpdateCertificate updates Certificate
func (cli *ZSClient) UpdateCertificate(ctx context.Context, uuid string, params param.UpdateCertificateParam) (*view.CertificateInventoryView, error) {
	resp := view.CertificateInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/certificates", uuid, "actions", "inventory", map[string]interface{}{
		"updateCertificate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateCertificate creates Certificate
func (cli *ZSClient) CreateCertificate(ctx context.Context, params param.CreateCertificateParam) (*view.CertificateInventoryView, error) {
	resp := view.CertificateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/certificates"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryCertificate queries Certificate list
func (cli *ZSClient) QueryCertificate(ctx context.Context, params *param.QueryParam) ([]view.CertificateInventoryView, error) {
	var resp []view.CertificateInventoryView
	return resp, cli.List(ctx, "v1/certificates", params, &resp)
}

func (cli *ZSClient) GetCertificate(ctx context.Context, uuid string) (*view.CertificateInventoryView, error) {
	var resp view.CertificateInventoryView
	if err := cli.Get(ctx, "v1/certificates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCertificate Pagination
func (cli *ZSClient) PageCertificate(ctx context.Context, params *param.QueryParam) ([]view.CertificateInventoryView, int, error) {
	var certificates []view.CertificateInventoryView
	total, err := cli.Page(ctx, "v1/certificates", params, &certificates)
	return certificates, total, err
}
