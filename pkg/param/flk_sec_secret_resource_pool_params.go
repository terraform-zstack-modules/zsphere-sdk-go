// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateFlkSecSecretResourcePoolParamDetail CreateFlkSecSecretResourcePool detail param
type CreateFlkSecSecretResourcePoolParamDetail struct {
	TestResult *string `json:"testResult,omitempty"`
	ActivatedToken *string `json:"activatedToken,omitempty"`
	ProtectToken *string `json:"protectToken,omitempty"`
	HmacToken *string `json:"hmacToken,omitempty"`
	UkeyType *string `json:"ukeyType,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Model *string `json:"model,omitempty"`
	Type string `json:"type" validate:"required"`
	HeartbeatInterval int `json:"heartbeatInterval" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateFlkSecSecretResourcePoolParam CreateFlkSecSecretResourcePool request param
type CreateFlkSecSecretResourcePoolParam struct {
	BaseParam
	Params CreateFlkSecSecretResourcePoolParamDetail `json:"params"`
}
// UpdateFlkSecSecretResourcePoolParamDetail UpdateFlkSecSecretResourcePool detail param
type UpdateFlkSecSecretResourcePoolParamDetail struct {
	EncryptResult *string `json:"encryptResult,omitempty"`
	ActivatedToken *string `json:"activatedToken,omitempty"`
	ProtectToken *string `json:"protectToken,omitempty"`
	HmacToken *string `json:"hmacToken,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Model *string `json:"model,omitempty"`
	HeartbeatInterval *int `json:"heartbeatInterval,omitempty"`
}

// UpdateFlkSecSecretResourcePoolParam UpdateFlkSecSecretResourcePool request param
type UpdateFlkSecSecretResourcePoolParam struct {
	BaseParam
	Params UpdateFlkSecSecretResourcePoolParamDetail `json:"updateFlkSecSecretResourcePool"`
}
