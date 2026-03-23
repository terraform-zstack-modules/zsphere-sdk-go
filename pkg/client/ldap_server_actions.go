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

// AddLdapServer adds LdapServer
func (cli *ZSClient) AddLdapServer(ctx context.Context, params param.AddLdapServerParam) (*view.LdapServerInventoryView, error) {
	resp := view.LdapServerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/ldap/servers"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteLdapServer deletes LdapServer
func (cli *ZSClient) DeleteLdapServer(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ldap/servers", uuid, string(deleteMode))
}
// QueryLdapServer queries LdapServer list
func (cli *ZSClient) QueryLdapServer(ctx context.Context, params *param.QueryParam) ([]view.LdapServerInventoryView, error) {
	var resp []view.LdapServerInventoryView
	return resp, cli.List(ctx, "v1/ldap/servers", params, &resp)
}

func (cli *ZSClient) GetLdapServer(ctx context.Context, uuid string) (*view.LdapServerInventoryView, error) {
	var resp view.LdapServerInventoryView
	if err := cli.Get(ctx, "v1/ldap/servers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLdapServer Pagination
func (cli *ZSClient) PageLdapServer(ctx context.Context, params *param.QueryParam) ([]view.LdapServerInventoryView, int, error) {
	var ldapServers []view.LdapServerInventoryView
	total, err := cli.Page(ctx, "v1/ldap/servers", params, &ldapServers)
	return ldapServers, total, err
}
// UpdateLdapServer updates LdapServer
func (cli *ZSClient) UpdateLdapServer(ctx context.Context, ldapServerUuid string, params param.UpdateLdapServerParam) (*view.LdapServerInventoryView, error) {
	resp := view.LdapServerInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/ldap/servers", ldapServerUuid, "", "inventory", map[string]interface{}{
		"updateLdapServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
