// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeletePortGroupParamDetail DeletePortGroup detail param
type DeletePortGroupParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeletePortGroupParam DeletePortGroup request param
type DeletePortGroupParam struct {
	BaseParam
	Params DeletePortGroupParamDetail `json:"deletePortGroup"`
}
// CreatePortGroupParamDetail CreatePortGroup detail param
type CreatePortGroupParamDetail struct {
	VSwitchUuid string `json:"vSwitchUuid" validate:"required"`
	VlanMode *string `json:"vlanMode,omitempty"`
	Vlan int `json:"vlan" validate:"required"`
	VlanRanges *string `json:"vlanRanges,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	L2NetworkUuid *string `json:"l2NetworkUuid,omitempty"`
	Category *string `json:"category,omitempty"`
	IpVersion *int `json:"ipVersion,omitempty"`
	System bool `json:"system,omitempty"`
	DnsDomain *string `json:"dnsDomain,omitempty"`
	EnableIPAM *bool `json:"enableIPAM,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePortGroupParam CreatePortGroup request param
type CreatePortGroupParam struct {
	BaseParam
	Params CreatePortGroupParamDetail `json:"params"`
}
// UpdatePortGroupParamDetail UpdatePortGroup detail param
type UpdatePortGroupParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	DnsDomain *string `json:"dnsDomain,omitempty"`
	Category *string `json:"category,omitempty"`
	System *bool `json:"system,omitempty"`
}

// UpdatePortGroupParam UpdatePortGroup request param
type UpdatePortGroupParam struct {
	BaseParam
	Params UpdatePortGroupParamDetail `json:"updatePortGroup"`
}
