// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateInfoSecSecretResourcePoolParamDetail CreateInfoSecSecretResourcePool detail param
type CreateInfoSecSecretResourcePoolParamDetail struct {
	ConnectionMode int `json:"connectionMode" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Model *string `json:"model,omitempty"`
	Type string `json:"type" validate:"required"`
	HeartbeatInterval int `json:"heartbeatInterval" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateInfoSecSecretResourcePoolParam CreateInfoSecSecretResourcePool request param
type CreateInfoSecSecretResourcePoolParam struct {
	BaseParam
	Params CreateInfoSecSecretResourcePoolParamDetail `json:"params"`
}
// UpdateInfoSecSecretResourcePoolParamDetail UpdateInfoSecSecretResourcePool detail param
type UpdateInfoSecSecretResourcePoolParamDetail struct {
	ConnectionMode *int `json:"connectionMode,omitempty"`
	ActivatedToken *string `json:"activatedToken,omitempty"`
	ProtectToken *string `json:"protectToken,omitempty"`
	HmacToken *string `json:"hmacToken,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Model *string `json:"model,omitempty"`
	HeartbeatInterval *int `json:"heartbeatInterval,omitempty"`
}

// UpdateInfoSecSecretResourcePoolParam UpdateInfoSecSecretResourcePool request param
type UpdateInfoSecSecretResourcePoolParam struct {
	BaseParam
	Params UpdateInfoSecSecretResourcePoolParamDetail `json:"updateInfoSecSecretResourcePool"`
}
