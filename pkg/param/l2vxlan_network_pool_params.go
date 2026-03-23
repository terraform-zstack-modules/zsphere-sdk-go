// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateL2VxlanNetworkPoolParamDetail CreateL2VxlanNetworkPool detail param
type CreateL2VxlanNetworkPoolParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	PhysicalInterface *string `json:"physicalInterface,omitempty"`
	Type *string `json:"type,omitempty"`
	VSwitchType *string `json:"vSwitchType,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL2VxlanNetworkPoolParam CreateL2VxlanNetworkPool request param
type CreateL2VxlanNetworkPoolParam struct {
	BaseParam
	Params CreateL2VxlanNetworkPoolParamDetail `json:"params"`
}
