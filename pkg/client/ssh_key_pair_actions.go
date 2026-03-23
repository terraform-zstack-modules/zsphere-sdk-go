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

// CreateSshKeyPair creates SshKeyPair
func (cli *ZSClient) CreateSshKeyPair(ctx context.Context, params param.CreateSshKeyPairParam) (*view.SshKeyPairInventoryView, error) {
	resp := view.SshKeyPairInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/ssh-key-pair"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// GenerateSshKeyPair operates on SshKeyPair
func (cli *ZSClient) GenerateSshKeyPair(ctx context.Context, params param.GenerateSshKeyPairParam) (*view.SshPrivateKeyPairInventoryView, error) {
	resp := view.SshPrivateKeyPairInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/ssh-key-pair/generate"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSshKeyPair updates SshKeyPair
func (cli *ZSClient) UpdateSshKeyPair(ctx context.Context, uuid string, params param.UpdateSshKeyPairParam) (*view.SshKeyPairInventoryView, error) {
	resp := view.SshKeyPairInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/ssh-key-pair", uuid, "actions", "inventory", map[string]interface{}{
		"updateSshKeyPair": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySshKeyPair queries SshKeyPair list
func (cli *ZSClient) QuerySshKeyPair(ctx context.Context, params *param.QueryParam) ([]view.SshKeyPairInventoryView, error) {
	var resp []view.SshKeyPairInventoryView
	return resp, cli.List(ctx, "v1/ssh-key-pair", params, &resp)
}

func (cli *ZSClient) GetSshKeyPair(ctx context.Context, uuid string) (*view.SshKeyPairInventoryView, error) {
	var resp view.SshKeyPairInventoryView
	if err := cli.Get(ctx, "v1/ssh-key-pair", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSshKeyPair Pagination
func (cli *ZSClient) PageSshKeyPair(ctx context.Context, params *param.QueryParam) ([]view.SshKeyPairInventoryView, int, error) {
	var sshKeyPairs []view.SshKeyPairInventoryView
	total, err := cli.Page(ctx, "v1/ssh-key-pair", params, &sshKeyPairs)
	return sshKeyPairs, total, err
}
// DeleteSshKeyPair deletes SshKeyPair
func (cli *ZSClient) DeleteSshKeyPair(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ssh-key-pair", uuid, string(deleteMode))
}
