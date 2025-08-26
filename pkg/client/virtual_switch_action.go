package client

import (
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

// QueryL2VirtualSwitchNetwork queries Layer 2 networks
func (cli *ZSClient) QueryL2VirtualSwitchNetwork(params param.QueryParam) ([]view.L2VirtualSwitchNetworkInventoryView, error) {
	resp := make([]view.L2VirtualSwitchNetworkInventoryView, 0)
	return resp, cli.List("v1/l2-networks/virtual-switch", &params, &resp)
}

// PageL2VirtualSwitchNetwork queries Layer 2 networks with pagination
func (cli *ZSClient) PageL2VirtualSwitchNetwork(params param.QueryParam) ([]view.L2VirtualSwitchNetworkInventoryView, int, error) {
	resp := make([]view.L2VirtualSwitchNetworkInventoryView, 0)
	total, err := cli.Page("v1/l2-networks/virtual-switch", &params, &resp)
	return resp, total, err
}

// GetL2VirtualSwitchNetwork queries a specific Layer 2 network
func (cli *ZSClient) GetL2VirtualSwitchNetwork(uuid string) (view.L2VirtualSwitchNetworkInventoryView, error) {
	resp := view.L2VirtualSwitchNetworkInventoryView{}
	return resp, cli.Get("v1/l2-networks/virtual-switch", uuid, nil, &resp)
}

// UpdateL2Network updates a Layer 2 network
func (cli *ZSClient) UpdateL2VirtualSwitchNetwork(uuid string, params param.UpdateL2NetworkParam) (view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	return resp, cli.Put("v1/l2-networks", uuid, &params, &resp)
}

// DeleteL2Network deletes a Layer 2 network
func (cli *ZSClient) DeleteL2VirtualSwitchNetwork(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks", uuid, string(deleteMode))
}

// CreateL2VirtualSwitch creates a virtual switch
func (cli *ZSClient) CreateL2VirtualSwitch(params param.CreateL2VirtualSwitchParam) (view.L2VirtualSwitchNetworkInventoryView, error) {
	resp := view.L2VirtualSwitchNetworkInventoryView{}
	return resp, cli.Post("v1/l2-networks/virtual-switch", &params, &resp)
}
