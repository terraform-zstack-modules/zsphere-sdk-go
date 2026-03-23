// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateInfoSecSecurityMachineParamDetail UpdateInfoSecSecurityMachine detail param
type UpdateInfoSecSecurityMachineParamDetail struct {
	Password *string `json:"password,omitempty"`
	Port *int `json:"port,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ManagementIp *string `json:"managementIp,omitempty"`
	Model *string `json:"model,omitempty"`
}

// UpdateInfoSecSecurityMachineParam UpdateInfoSecSecurityMachine request param
type UpdateInfoSecSecurityMachineParam struct {
	BaseParam
	Params UpdateInfoSecSecurityMachineParamDetail `json:"updateInfoSecSecurityMachine"`
}
// AddInfoSecSecurityMachineParamDetail AddInfoSecSecurityMachine detail param
type AddInfoSecSecurityMachineParamDetail struct {
	Password string `json:"password" validate:"required"`
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

// AddInfoSecSecurityMachineParam AddInfoSecSecurityMachine request param
type AddInfoSecSecurityMachineParam struct {
	BaseParam
	Params AddInfoSecSecurityMachineParamDetail `json:"params"`
}
