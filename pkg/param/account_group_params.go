// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateAccountGroupParamDetail CreateAccountGroup detail param
type CreateAccountGroupParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ParentUuid *string `json:"parentUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAccountGroupParam CreateAccountGroup request param
type CreateAccountGroupParam struct {
	BaseParam
	Params CreateAccountGroupParamDetail `json:"params"`
}
// UpdateAccountGroupParamDetail UpdateAccountGroup detail param
type UpdateAccountGroupParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateAccountGroupParam UpdateAccountGroup request param
type UpdateAccountGroupParam struct {
	BaseParam
	Params UpdateAccountGroupParamDetail `json:"updateAccountGroup"`
}
// DeleteAccountGroupParamDetail DeleteAccountGroup detail param
type DeleteAccountGroupParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAccountGroupParam DeleteAccountGroup request param
type DeleteAccountGroupParam struct {
	BaseParam
	Params DeleteAccountGroupParamDetail `json:"deleteAccountGroup"`
}
