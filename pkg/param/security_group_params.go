// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateSecurityGroupParamDetail CreateSecurityGroup detail param
type CreateSecurityGroupParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	IpVersion *int `json:"ipVersion,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSecurityGroupParam CreateSecurityGroup request param
type CreateSecurityGroupParam struct {
	BaseParam
	Params CreateSecurityGroupParamDetail `json:"params"`
}
// DeleteSecurityGroupParamDetail DeleteSecurityGroup detail param
type DeleteSecurityGroupParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteSecurityGroupParam DeleteSecurityGroup request param
type DeleteSecurityGroupParam struct {
	BaseParam
	Params DeleteSecurityGroupParamDetail `json:"deleteSecurityGroup"`
}
// UpdateSecurityGroupParamDetail UpdateSecurityGroup detail param
type UpdateSecurityGroupParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateSecurityGroupParam UpdateSecurityGroup request param
type UpdateSecurityGroupParam struct {
	BaseParam
	Params UpdateSecurityGroupParamDetail `json:"updateSecurityGroup"`
}
