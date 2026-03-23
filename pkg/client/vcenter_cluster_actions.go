// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVCenterCluster queries VCenterCluster list
func (cli *ZSClient) QueryVCenterCluster(ctx context.Context, params *param.QueryParam) ([]view.VCenterClusterInventoryView, error) {
	var resp []view.VCenterClusterInventoryView
	return resp, cli.List(ctx, "v1/vcenters/clusters", params, &resp)
}

func (cli *ZSClient) GetVCenterCluster(ctx context.Context, uuid string) (*view.VCenterClusterInventoryView, error) {
	var resp view.VCenterClusterInventoryView
	if err := cli.Get(ctx, "v1/vcenters/clusters", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVCenterCluster Pagination
func (cli *ZSClient) PageVCenterCluster(ctx context.Context, params *param.QueryParam) ([]view.VCenterClusterInventoryView, int, error) {
	var vCenterClusters []view.VCenterClusterInventoryView
	total, err := cli.Page(ctx, "v1/vcenters/clusters", params, &vCenterClusters)
	return vCenterClusters, total, err
}
