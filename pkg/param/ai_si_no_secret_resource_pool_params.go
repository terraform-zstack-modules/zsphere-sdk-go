// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateAiSiNoSecretResourcePoolParamDetail CreateAiSiNoSecretResourcePool detail param
type CreateAiSiNoSecretResourcePoolParamDetail struct {
	ManagementIp *string `json:"managementIp,omitempty"`
	Port *int `json:"port,omitempty"`
	Route *string `json:"route,omitempty"`
	ClientID *string `json:"clientID,omitempty"`
	ClientSecrete *string `json:"clientSecrete,omitempty"`
	AppId *string `json:"appId,omitempty"`
	KeyNumSM2 *string `json:"keyNumSM2,omitempty"`
	KeyNumSM4 *string `json:"keyNumSM4,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Model *string `json:"model,omitempty"`
	Type string `json:"type" validate:"required"`
	HeartbeatInterval int `json:"heartbeatInterval" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAiSiNoSecretResourcePoolParam CreateAiSiNoSecretResourcePool request param
type CreateAiSiNoSecretResourcePoolParam struct {
	BaseParam
	Params CreateAiSiNoSecretResourcePoolParamDetail `json:"params"`
}
