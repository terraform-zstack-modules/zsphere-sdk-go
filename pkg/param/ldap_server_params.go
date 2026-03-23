// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// AddLdapServerParamDetail AddLdapServer detail param
type AddLdapServerParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Url string `json:"url" validate:"required"`
	Base string `json:"base" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Encryption string `json:"encryption" validate:"required"`
	ServerType *string `json:"serverType,omitempty"`
	UsernameProperty *string `json:"usernameProperty,omitempty"`
	Filter *string `json:"filter,omitempty"`
	SyncCreatedAccountStrategy *string `json:"syncCreatedAccountStrategy,omitempty"`
	SyncDeletedAccountStrategy *string `json:"syncDeletedAccountStrategy,omitempty"`
}

// AddLdapServerParam AddLdapServer request param
type AddLdapServerParam struct {
	BaseParam
	Params AddLdapServerParamDetail `json:"params"`
}
// DeleteLdapServerParamDetail DeleteLdapServer detail param
type DeleteLdapServerParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteLdapServerParam DeleteLdapServer request param
type DeleteLdapServerParam struct {
	BaseParam
	Params DeleteLdapServerParamDetail `json:"deleteLdapServer"`
}
// UpdateLdapServerParamDetail UpdateLdapServer detail param
type UpdateLdapServerParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Url *string `json:"url,omitempty"`
	Base *string `json:"base,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Encryption *string `json:"encryption,omitempty"`
	ServerType *string `json:"serverType,omitempty"`
	UsernameProperty *string `json:"usernameProperty,omitempty"`
	Filter *string `json:"filter,omitempty"`
	SyncCreatedAccountStrategy *string `json:"syncCreatedAccountStrategy,omitempty"`
	SyncDeletedAccountStrategy *string `json:"syncDeletedAccountStrategy,omitempty"`
}

// UpdateLdapServerParam UpdateLdapServer request param
type UpdateLdapServerParam struct {
	BaseParam
	Params UpdateLdapServerParamDetail `json:"updateLdapServer"`
}
