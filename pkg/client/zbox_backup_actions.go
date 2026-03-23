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

// QueryZBoxBackup queries ZBoxBackup list
func (cli *ZSClient) QueryZBoxBackup(ctx context.Context, params *param.QueryParam) ([]view.ZBoxBackupInventoryView, error) {
	var resp []view.ZBoxBackupInventoryView
	return resp, cli.List(ctx, "v1/externalbackup/zbox", params, &resp)
}

func (cli *ZSClient) GetZBoxBackup(ctx context.Context, uuid string) (*view.ZBoxBackupInventoryView, error) {
	var resp view.ZBoxBackupInventoryView
	if err := cli.Get(ctx, "v1/externalbackup/zbox", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageZBoxBackup Pagination
func (cli *ZSClient) PageZBoxBackup(ctx context.Context, params *param.QueryParam) ([]view.ZBoxBackupInventoryView, int, error) {
	var zBoxBackups []view.ZBoxBackupInventoryView
	total, err := cli.Page(ctx, "v1/externalbackup/zbox", params, &zBoxBackups)
	return zBoxBackups, total, err
}
// CreateZBoxBackup creates ZBoxBackup
func (cli *ZSClient) CreateZBoxBackup(ctx context.Context, params param.CreateZBoxBackupParam) (*view.ExternalBackupInventoryView, error) {
	resp := view.ExternalBackupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/externalbackup/zbox"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateZBoxBackupAsync Async
func (cli *ZSClient) CreateZBoxBackupAsync(ctx context.Context, params param.CreateZBoxBackupParam) (string, error) {

	resource := "externalbackup/zbox"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
