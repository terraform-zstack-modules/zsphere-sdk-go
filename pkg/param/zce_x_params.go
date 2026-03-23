// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// RemoveZceXParamDetail RemoveZceX detail param
type RemoveZceXParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RemoveZceXParam RemoveZceX request param
type RemoveZceXParam struct {
	BaseParam
	Params RemoveZceXParamDetail `json:"removeZceX"`
}
// AddZceXParamDetail AddZceX detail param
type AddZceXParamDetail struct {
	Name string `json:"name" validate:"required"`
	AdminToken *string `json:"adminToken,omitempty"`
	ManagementIp *string `json:"managementIp,omitempty"`
	ApiPort *int `json:"apiPort,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddZceXParam AddZceX request param
type AddZceXParam struct {
	BaseParam
	Params AddZceXParamDetail `json:"params"`
}
