// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateL2VxlanNetworkParamDetail CreateL2VxlanNetwork detail param
type CreateL2VxlanNetworkParamDetail struct {
	Vni *int `json:"vni,omitempty"`
	PoolUuid string `json:"poolUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	PhysicalInterface *string `json:"physicalInterface,omitempty"`
	Type *string `json:"type,omitempty"`
	VSwitchType *string `json:"vSwitchType,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL2VxlanNetworkParam CreateL2VxlanNetwork request param
type CreateL2VxlanNetworkParam struct {
	BaseParam
	Params CreateL2VxlanNetworkParamDetail `json:"params"`
}
