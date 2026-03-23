// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateSnmpAgentParamDetail CreateSnmpAgent detail param
type CreateSnmpAgentParamDetail struct {
	Version string `json:"version" validate:"required"`
	ReadCommunity *string `json:"readCommunity,omitempty"`
	UserName *string `json:"userName,omitempty"`
	AuthAlgorithm *string `json:"authAlgorithm,omitempty"`
	AuthPassword *string `json:"authPassword,omitempty"`
	PrivacyAlgorithm *string `json:"privacyAlgorithm,omitempty"`
	PrivacyPassword *string `json:"privacyPassword,omitempty"`
	Port int `json:"port" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSnmpAgentParam CreateSnmpAgent request param
type CreateSnmpAgentParam struct {
	BaseParam
	Params CreateSnmpAgentParamDetail `json:"params"`
}
// StartSnmpAgentParamDetail StartSnmpAgent detail param
type StartSnmpAgentParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// StartSnmpAgentParam StartSnmpAgent request param
type StartSnmpAgentParam struct {
	BaseParam
	Params StartSnmpAgentParamDetail `json:"startSnmpAgent"`
}
// StopSnmpAgentParamDetail StopSnmpAgent detail param
type StopSnmpAgentParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// StopSnmpAgentParam StopSnmpAgent request param
type StopSnmpAgentParam struct {
	BaseParam
	Params StopSnmpAgentParamDetail `json:"stopSnmpAgent"`
}
// UpdateSnmpAgentParamDetail UpdateSnmpAgent detail param
type UpdateSnmpAgentParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Version string `json:"version" validate:"required"`
	ReadCommunity *string `json:"readCommunity,omitempty"`
	UserName *string `json:"userName,omitempty"`
	AuthAlgorithm *string `json:"authAlgorithm,omitempty"`
	AuthPassword *string `json:"authPassword,omitempty"`
	PrivacyAlgorithm *string `json:"privacyAlgorithm,omitempty"`
	PrivacyPassword *string `json:"privacyPassword,omitempty"`
	Port int `json:"port" validate:"required"`
}

// UpdateSnmpAgentParam UpdateSnmpAgent request param
type UpdateSnmpAgentParam struct {
	BaseParam
	Params UpdateSnmpAgentParamDetail `json:"updateSnmpAgent"`
}
