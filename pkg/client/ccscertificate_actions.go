// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// AddCCSCertificate adds CCSCertificate
func (cli *ZSClient) AddCCSCertificate(ctx context.Context, params param.AddCCSCertificateParam) (*view.CCSCertificateInventoryView, error) {
	resp := view.CCSCertificateInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/crypto/ccs-certificate/add", "", "", map[string]interface{}{
		"addCCSCertificate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteCCSCertificate deletes CCSCertificate
func (cli *ZSClient) DeleteCCSCertificate(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/crypto/ccs-certificate/delete", uuid, string(deleteMode))
}
// QueryCCSCertificate queries CCSCertificate list
func (cli *ZSClient) QueryCCSCertificate(ctx context.Context, params *param.QueryParam) ([]view.CCSCertificateInventoryView, error) {
	var resp []view.CCSCertificateInventoryView
	return resp, cli.List(ctx, "v1/crypto/ccs-certificate/certificates/", params, &resp)
}

func (cli *ZSClient) GetCCSCertificate(ctx context.Context, uuid string) (*view.CCSCertificateInventoryView, error) {
	var resp view.CCSCertificateInventoryView
	if err := cli.Get(ctx, "v1/crypto/ccs-certificate/certificates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCCSCertificate Pagination
func (cli *ZSClient) PageCCSCertificate(ctx context.Context, params *param.QueryParam) ([]view.CCSCertificateInventoryView, int, error) {
	var cCSCertificates []view.CCSCertificateInventoryView
	total, err := cli.Page(ctx, "v1/crypto/ccs-certificate/certificates/", params, &cCSCertificates)
	return cCSCertificates, total, err
}
