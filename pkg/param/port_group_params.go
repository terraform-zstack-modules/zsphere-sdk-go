// Copyright (c) ZStack.io, Inc.

package param

type UpdatePortGroupParam struct {
	BaseParam
	UpdateL3Network UpdatePortGroupDetailParam `json:"updatePortGroup"`
}

type UpdatePortGroupDetailParam struct {
	BaseParam
	Name        string      `json:"name"`        //  port group name
	Description *string     `json:"description"` // port group description
	System      *bool       `json:"system"`      // Whether it is for system cloud hosts
	DnsDomain   *string     `json:"dnsDomain"`   // DNS domain of the port group
	Category    *L3Category `json:"category"`    // Category of the port group
}

type CreatePortGroupParam struct {
	BaseParam
	Params CreatePortGroupDetailParam `json:"params"`
}

type CreatePortGroupDetailParam struct {
	VSwitchUuid string `json:"vSwitchUuid"`
	VlanMode    string `json:"vlanMode"` //ACCESS, TRUNK, PVLAN
	Vlan        int16  `json:"vlan"`
	VlanRanges  string `json:"vlanRanges"`

	Name        string `json:"name"`
	Description string `json:"description"` // port group description
	Type        string `json:"type"`
	Category    string `json:"category"`
	IpVersion   int    `json:"ipVersion"` // 4, 6
	System      bool   `json:"system"`
	DnsDomain   string `json:"dnsDomain"`
	EnableIPAM  bool   `json:"enableIPAM"`
}
