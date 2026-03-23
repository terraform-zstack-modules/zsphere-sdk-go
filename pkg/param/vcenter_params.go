// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// AddVCenterParamDetail AddVCenter detail param
type AddVCenterParamDetail struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Https *bool `json:"https,omitempty"`
	Port *int `json:"port,omitempty"`
	DomainName string `json:"domainName" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddVCenterParam AddVCenter request param
type AddVCenterParam struct {
	BaseParam
	Params AddVCenterParamDetail `json:"params"`
}
// SyncVCenterParamDetail SyncVCenter detail param
type SyncVCenterParamDetail struct {
}

// SyncVCenterParam SyncVCenter request param
type SyncVCenterParam struct {
	BaseParam
	Params SyncVCenterParamDetail `json:"syncVCenter"`
}
// UpdateVCenterParamDetail UpdateVCenter detail param
type UpdateVCenterParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	DomainName *string `json:"domainName,omitempty"`
	Port *int `json:"port,omitempty"`
	State *string `json:"state,omitempty"`
}

// UpdateVCenterParam UpdateVCenter request param
type UpdateVCenterParam struct {
	BaseParam
	Params UpdateVCenterParamDetail `json:"updateVCenter"`
}
// DeleteVCenterParamDetail DeleteVCenter detail param
type DeleteVCenterParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVCenterParam DeleteVCenter request param
type DeleteVCenterParam struct {
	BaseParam
	Params DeleteVCenterParamDetail `json:"deleteVCenter"`
}
