// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteNvmeServerParamDetail DeleteNvmeServer detail param
type DeleteNvmeServerParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteNvmeServerParam DeleteNvmeServer request param
type DeleteNvmeServerParam struct {
	BaseParam
	Params DeleteNvmeServerParamDetail `json:"deleteNvmeServer"`
}
// RefreshNvmeServerParamDetail RefreshNvmeServer detail param
type RefreshNvmeServerParamDetail struct {
}

// RefreshNvmeServerParam RefreshNvmeServer request param
type RefreshNvmeServerParam struct {
	BaseParam
	Params RefreshNvmeServerParamDetail `json:"refreshNvmeServer"`
}
// UpdateNvmeServerParamDetail UpdateNvmeServer detail param
type UpdateNvmeServerParamDetail struct {
	Name string `json:"name,omitempty"`
}

// UpdateNvmeServerParam UpdateNvmeServer request param
type UpdateNvmeServerParam struct {
	BaseParam
	Params UpdateNvmeServerParamDetail `json:"updateNvmeServer"`
}
// AddNvmeServerParamDetail AddNvmeServer detail param
type AddNvmeServerParamDetail struct {
	Name string `json:"name,omitempty"`
	Ip string `json:"ip" validate:"required"`
	Port *int `json:"port,omitempty"`
	Transport string `json:"transport" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddNvmeServerParam AddNvmeServer request param
type AddNvmeServerParam struct {
	BaseParam
	Params AddNvmeServerParamDetail `json:"params"`
}
