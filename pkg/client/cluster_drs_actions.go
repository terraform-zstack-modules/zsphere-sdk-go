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

// CreateClusterDRS creates ClusterDRS
func (cli *ZSClient) CreateClusterDRS(ctx context.Context, clusterUuid string, params param.CreateClusterDRSParam) (*view.ClusterDRSInventoryView, error) {
	resp := view.ClusterDRSInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/clusters/%s/drs", clusterUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateClusterDRS updates ClusterDRS
func (cli *ZSClient) UpdateClusterDRS(ctx context.Context, uuid string, params param.UpdateClusterDRSParam) (*view.ClusterDRSInventoryView, error) {
	resp := view.ClusterDRSInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/clusters/drs", uuid, "actions", "inventory", map[string]interface{}{
		"updateClusterDRS": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryClusterDRS queries ClusterDRS list
func (cli *ZSClient) QueryClusterDRS(ctx context.Context, params *param.QueryParam) ([]view.ClusterDRSInventoryView, error) {
	var resp []view.ClusterDRSInventoryView
	return resp, cli.List(ctx, "v1/clusters/drs", params, &resp)
}

func (cli *ZSClient) GetClusterDRS(ctx context.Context, uuid string) (*view.ClusterDRSInventoryView, error) {
	var resp view.ClusterDRSInventoryView
	if err := cli.Get(ctx, "v1/clusters/drs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageClusterDRS Pagination
func (cli *ZSClient) PageClusterDRS(ctx context.Context, params *param.QueryParam) ([]view.ClusterDRSInventoryView, int, error) {
	var clusterDRSs []view.ClusterDRSInventoryView
	total, err := cli.Page(ctx, "v1/clusters/drs", params, &clusterDRSs)
	return clusterDRSs, total, err
}
// DeleteClusterDRS deletes ClusterDRS
func (cli *ZSClient) DeleteClusterDRS(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/clusters/drs", uuid, string(deleteMode))
}
