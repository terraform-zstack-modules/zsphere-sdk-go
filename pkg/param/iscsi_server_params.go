// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// RefreshIscsiServerParamDetail RefreshIscsiServer detail param
type RefreshIscsiServerParamDetail struct {
}

// RefreshIscsiServerParam RefreshIscsiServer request param
type RefreshIscsiServerParam struct {
	BaseParam
	Params RefreshIscsiServerParamDetail `json:"params"`
}
// AddIscsiServerParamDetail AddIscsiServer detail param
type AddIscsiServerParamDetail struct {
	Name string `json:"name,omitempty"`
	Ip string `json:"ip" validate:"required"`
	Port *int `json:"port,omitempty"`
	ChapUserName *string `json:"chapUserName,omitempty"`
	ChapUserPassword *string `json:"chapUserPassword,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddIscsiServerParam AddIscsiServer request param
type AddIscsiServerParam struct {
	BaseParam
	Params AddIscsiServerParamDetail `json:"params"`
}
// UpdateIscsiServerParamDetail UpdateIscsiServer detail param
type UpdateIscsiServerParamDetail struct {
	Name string `json:"name,omitempty"`
	ChapUserName *string `json:"chapUserName,omitempty"`
	ChapUserPassword *string `json:"chapUserPassword,omitempty"`
	State *string `json:"state,omitempty"`
}

// UpdateIscsiServerParam UpdateIscsiServer request param
type UpdateIscsiServerParam struct {
	BaseParam
	Params UpdateIscsiServerParamDetail `json:"updateIscsiServer"`
}
// DeleteIscsiServerParamDetail DeleteIscsiServer detail param
type DeleteIscsiServerParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteIscsiServerParam DeleteIscsiServer request param
type DeleteIscsiServerParam struct {
	BaseParam
	Params DeleteIscsiServerParamDetail `json:"deleteIscsiServer"`
}
