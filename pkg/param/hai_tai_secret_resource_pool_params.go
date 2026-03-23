// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateHaiTaiSecretResourcePoolParamDetail CreateHaiTaiSecretResourcePool detail param
type CreateHaiTaiSecretResourcePoolParamDetail struct {
	ManagementIp *string `json:"managementIp,omitempty"`
	Port *int `json:"port,omitempty"`
	Realm *string `json:"realm,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Model *string `json:"model,omitempty"`
	Type string `json:"type" validate:"required"`
	HeartbeatInterval int `json:"heartbeatInterval" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateHaiTaiSecretResourcePoolParam CreateHaiTaiSecretResourcePool request param
type CreateHaiTaiSecretResourcePoolParam struct {
	BaseParam
	Params CreateHaiTaiSecretResourcePoolParamDetail `json:"params"`
}
