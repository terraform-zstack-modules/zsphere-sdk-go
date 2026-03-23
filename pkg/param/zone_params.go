// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteZoneParamDetail DeleteZone detail param
type DeleteZoneParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteZoneParam DeleteZone request param
type DeleteZoneParam struct {
	BaseParam
	Params DeleteZoneParamDetail `json:"deleteZone"`
}
// UpdateZoneParamDetail UpdateZone detail param
type UpdateZoneParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
}

// UpdateZoneParam UpdateZone request param
type UpdateZoneParam struct {
	BaseParam
	Params UpdateZoneParamDetail `json:"updateZone"`
}
// GetZoneParamDetail GetZone detail param
type GetZoneParamDetail struct {
}

// GetZoneParam GetZone request param
type GetZoneParam struct {
	BaseParam
	Params GetZoneParamDetail `json:"getZone"`
}
// CreateZoneParamDetail CreateZone detail param
type CreateZoneParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateZoneParam CreateZone request param
type CreateZoneParam struct {
	BaseParam
	Params CreateZoneParamDetail `json:"params"`
}
