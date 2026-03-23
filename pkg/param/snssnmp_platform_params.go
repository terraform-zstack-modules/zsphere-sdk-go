// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateSNSSnmpPlatformParamDetail CreateSNSSnmpPlatform detail param
type CreateSNSSnmpPlatformParamDetail struct {
	SnmpAddress string `json:"snmpAddress" validate:"required"`
	SnmpPort int `json:"snmpPort" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSSnmpPlatformParam CreateSNSSnmpPlatform request param
type CreateSNSSnmpPlatformParam struct {
	BaseParam
	Params CreateSNSSnmpPlatformParamDetail `json:"params"`
}
// UpdateSNSSnmpPlatformParamDetail UpdateSNSSnmpPlatform detail param
type UpdateSNSSnmpPlatformParamDetail struct {
	SnmpAddress string `json:"snmpAddress" validate:"required"`
	SnmpPort int `json:"snmpPort" validate:"required"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateSNSSnmpPlatformParam UpdateSNSSnmpPlatform request param
type UpdateSNSSnmpPlatformParam struct {
	BaseParam
	Params UpdateSNSSnmpPlatformParamDetail `json:"updateSNSSnmpPlatform"`
}
