// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryDRSVmMigrationActivity queries DRSVmMigrationActivity list
func (cli *ZSClient) QueryDRSVmMigrationActivity(ctx context.Context, params *param.QueryParam) ([]view.DRSVmMigrationActivityInventoryView, error) {
	var resp []view.DRSVmMigrationActivityInventoryView
	return resp, cli.List(ctx, "v1/clusters/drs/vm-migration-activities", params, &resp)
}

func (cli *ZSClient) GetDRSVmMigrationActivity(ctx context.Context, uuid string) (*view.DRSVmMigrationActivityInventoryView, error) {
	var resp view.DRSVmMigrationActivityInventoryView
	if err := cli.Get(ctx, "v1/clusters/drs/vm-migration-activities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageDRSVmMigrationActivity Pagination
func (cli *ZSClient) PageDRSVmMigrationActivity(ctx context.Context, params *param.QueryParam) ([]view.DRSVmMigrationActivityInventoryView, int, error) {
	var dRSVmMigrationActivities []view.DRSVmMigrationActivityInventoryView
	total, err := cli.Page(ctx, "v1/clusters/drs/vm-migration-activities", params, &dRSVmMigrationActivities)
	return dRSVmMigrationActivities, total, err
}
