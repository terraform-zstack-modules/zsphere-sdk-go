// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteSecurityMachineParamDetail DeleteSecurityMachine detail param
type DeleteSecurityMachineParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteSecurityMachineParam DeleteSecurityMachine request param
type DeleteSecurityMachineParam struct {
	BaseParam
	Params DeleteSecurityMachineParamDetail `json:"deleteSecurityMachine"`
}
// UpdateSecurityMachineParamDetail UpdateSecurityMachine detail param
type UpdateSecurityMachineParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ManagementIp *string `json:"managementIp,omitempty"`
	Model *string `json:"model,omitempty"`
}

// UpdateSecurityMachineParam UpdateSecurityMachine request param
type UpdateSecurityMachineParam struct {
	BaseParam
	Params UpdateSecurityMachineParamDetail `json:"updateSecurityMachine"`
}
