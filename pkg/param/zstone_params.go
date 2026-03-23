// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// AddZStoneParamDetail AddZStone detail param
type AddZStoneParamDetail struct {
	Name string `json:"name" validate:"required"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	ManagementIp string `json:"managementIp" validate:"required"`
	LogInPort *int `json:"logInPort,omitempty"`
	ApiPort *int `json:"apiPort,omitempty"`
	LogInUrl *string `json:"logInUrl,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddZStoneParam AddZStone request param
type AddZStoneParam struct {
	BaseParam
	Params AddZStoneParamDetail `json:"params"`
}
// RemoveZStoneParamDetail RemoveZStone detail param
type RemoveZStoneParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RemoveZStoneParam RemoveZStone request param
type RemoveZStoneParam struct {
	BaseParam
	Params RemoveZStoneParamDetail `json:"removeZStone"`
}
