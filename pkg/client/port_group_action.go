package client

import (
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

// QueryPortGroup Queries Layer Port Group
func (cli *ZSClient) QueryPortGroup(params param.QueryParam) ([]view.PortGroupInventoryView, error) {
	var network []view.PortGroupInventoryView
	return network, cli.List("v1/l3-networks/port-group", &params, &network)
}

// PagePortGroup Paginated query for Layer 3 networks
func (cli *ZSClient) PagePortGroup(params param.QueryParam) ([]view.PortGroupInventoryView, int, error) {
	var network []view.PortGroupInventoryView
	total, err := cli.Page("v1/l3-networks/port-group", &params, &network)
	return network, total, err
}

// GetPortGroup Queries a specific Layer 3 network
func (cli *ZSClient) GetPortGroup(uuid string) (view.PortGroupInventoryView, error) {
	var resp view.PortGroupInventoryView
	return resp, cli.Get("v1/l3-networks/port-group", uuid, nil, &resp)
}

// UpdateL3Network Updates a Layer 3 network
func (cli *ZSClient) UpdatePortGroup(uuid string, params param.UpdatePortGroupParam) (view.PortGroupInventoryView, error) {
	var resp view.PortGroupInventoryView
	return resp, cli.Put("v1/l3-networks/port-group", uuid, &params, &resp)
}

// DeleteL3Network Deletes a Layer 3 network
func (cli *ZSClient) DeletePortGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/port-group", uuid, string(deleteMode))
}

// CreateL3Network Creates a Layer 3 network
func (cli *ZSClient) CreatePortGroup(params param.CreatePortGroupParam) (view.PortGroupInventoryView, error) {
	var resp view.PortGroupInventoryView
	return resp, cli.Post("v1/l3-networks/port-group", &params, &resp)
}
