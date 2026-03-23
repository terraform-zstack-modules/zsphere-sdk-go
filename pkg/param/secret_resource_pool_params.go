// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateSecretResourcePoolParamDetail UpdateSecretResourcePool detail param
type UpdateSecretResourcePoolParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Model *string `json:"model,omitempty"`
	HeartbeatInterval *int `json:"heartbeatInterval,omitempty"`
}

// UpdateSecretResourcePoolParam UpdateSecretResourcePool request param
type UpdateSecretResourcePoolParam struct {
	BaseParam
	Params UpdateSecretResourcePoolParamDetail `json:"updateSecretResourcePool"`
}
// DeleteSecretResourcePoolParamDetail DeleteSecretResourcePool detail param
type DeleteSecretResourcePoolParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteSecretResourcePoolParam DeleteSecretResourcePool request param
type DeleteSecretResourcePoolParam struct {
	BaseParam
	Params DeleteSecretResourcePoolParamDetail `json:"deleteSecretResourcePool"`
}
