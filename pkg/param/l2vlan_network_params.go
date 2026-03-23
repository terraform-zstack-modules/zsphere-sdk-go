// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateL2VlanNetworkParamDetail CreateL2VlanNetwork detail param
type CreateL2VlanNetworkParamDetail struct {
	Vlan int `json:"vlan" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	PhysicalInterface string `json:"physicalInterface" validate:"required"`
	Type *string `json:"type,omitempty"`
	VSwitchType *string `json:"vSwitchType,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL2VlanNetworkParam CreateL2VlanNetwork request param
type CreateL2VlanNetworkParam struct {
	BaseParam
	Params CreateL2VlanNetworkParamDetail `json:"params"`
}
