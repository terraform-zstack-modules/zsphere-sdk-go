// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateSlbGroupParamDetail CreateSlbGroup detail param
type CreateSlbGroupParamDetail struct {
	Name string `json:"name" validate:"required"`
	SlbOfferingUuid string `json:"slbOfferingUuid" validate:"required"`
	FrontEndL3NetworkUuid string `json:"frontEndL3NetworkUuid" validate:"required"`
	BackendL3NetworkUuids []string `json:"backendL3NetworkUuids,omitempty"`
	BackendType *string `json:"backendType,omitempty"`
	DeployType *string `json:"deployType,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSlbGroupParam CreateSlbGroup request param
type CreateSlbGroupParam struct {
	BaseParam
	Params CreateSlbGroupParamDetail `json:"params"`
}
// DeleteSlbGroupParamDetail DeleteSlbGroup detail param
type DeleteSlbGroupParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteSlbGroupParam DeleteSlbGroup request param
type DeleteSlbGroupParam struct {
	BaseParam
	Params DeleteSlbGroupParamDetail `json:"deleteSlbGroup"`
}
// UpdateSlbGroupParamDetail UpdateSlbGroup detail param
type UpdateSlbGroupParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	SlbOfferingUuid *string `json:"slbOfferingUuid,omitempty"`
}

// UpdateSlbGroupParam UpdateSlbGroup request param
type UpdateSlbGroupParam struct {
	BaseParam
	Params UpdateSlbGroupParamDetail `json:"updateSlbGroup"`
}
