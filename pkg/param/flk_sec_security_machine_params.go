// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateFlkSecSecurityMachineParamDetail UpdateFlkSecSecurityMachine detail param
type UpdateFlkSecSecurityMachineParamDetail struct {
	Port *int `json:"port,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ManagementIp *string `json:"managementIp,omitempty"`
	Model *string `json:"model,omitempty"`
}

// UpdateFlkSecSecurityMachineParam UpdateFlkSecSecurityMachine request param
type UpdateFlkSecSecurityMachineParam struct {
	BaseParam
	Params UpdateFlkSecSecurityMachineParamDetail `json:"updateFlkSecSecurityMachine"`
}
// AddFlkSecSecurityMachineParamDetail AddFlkSecSecurityMachine detail param
type AddFlkSecSecurityMachineParamDetail struct {
	Port int `json:"port" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp" validate:"required"`
	Model string `json:"model" validate:"required"`
	Type string `json:"type" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	SecretResourcePoolUuid string `json:"secretResourcePoolUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddFlkSecSecurityMachineParam AddFlkSecSecurityMachine request param
type AddFlkSecSecurityMachineParam struct {
	BaseParam
	Params AddFlkSecSecurityMachineParamDetail `json:"params"`
}
