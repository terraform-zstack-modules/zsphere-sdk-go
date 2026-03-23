// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateVpcFirewallParamDetail UpdateVpcFirewall detail param
type UpdateVpcFirewallParamDetail struct {
	Description *string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
}

// UpdateVpcFirewallParam UpdateVpcFirewall request param
type UpdateVpcFirewallParam struct {
	BaseParam
	Params UpdateVpcFirewallParamDetail `json:"updateVpcFirewall"`
}
// CreateVpcFirewallParamDetail CreateVpcFirewall detail param
type CreateVpcFirewallParamDetail struct {
	VpcUuid string `json:"vpcUuid" validate:"required"`
	Description *string `json:"description,omitempty"`
	Name string `json:"name" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpcFirewallParam CreateVpcFirewall request param
type CreateVpcFirewallParam struct {
	BaseParam
	Params CreateVpcFirewallParamDetail `json:"params"`
}
