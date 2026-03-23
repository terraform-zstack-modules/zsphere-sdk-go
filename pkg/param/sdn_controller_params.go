// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// RemoveSdnControllerParamDetail RemoveSdnController detail param
type RemoveSdnControllerParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RemoveSdnControllerParam RemoveSdnController request param
type RemoveSdnControllerParam struct {
	BaseParam
	Params RemoveSdnControllerParamDetail `json:"removeSdnController"`
}
// AddSdnControllerParamDetail AddSdnController detail param
type AddSdnControllerParamDetail struct {
	VendorType string `json:"vendorType" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Ip string `json:"ip" validate:"required"`
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSdnControllerParam AddSdnController request param
type AddSdnControllerParam struct {
	BaseParam
	Params AddSdnControllerParamDetail `json:"params"`
}
// UpdateSdnControllerParamDetail UpdateSdnController detail param
type UpdateSdnControllerParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateSdnControllerParam UpdateSdnController request param
type UpdateSdnControllerParam struct {
	BaseParam
	Params UpdateSdnControllerParamDetail `json:"updateSdnController"`
}
