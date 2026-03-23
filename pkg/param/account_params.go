// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateAccountParamDetail UpdateAccount detail param
type UpdateAccountParamDetail struct {
	Password *string `json:"password,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	OldPassword *string `json:"oldPassword,omitempty"`
	State *string `json:"state,omitempty"`
}

// UpdateAccountParam UpdateAccount request param
type UpdateAccountParam struct {
	BaseParam
	Params UpdateAccountParamDetail `json:"updateAccount"`
}
// DeleteAccountParamDetail DeleteAccount detail param
type DeleteAccountParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAccountParam DeleteAccount request param
type DeleteAccountParam struct {
	BaseParam
	Params DeleteAccountParamDetail `json:"deleteAccount"`
}
// CreateAccountParamDetail CreateAccount detail param
type CreateAccountParamDetail struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Type *string `json:"type,omitempty"`
	State *string `json:"state,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAccountParam CreateAccount request param
type CreateAccountParam struct {
	BaseParam
	Params CreateAccountParamDetail `json:"params"`
}
