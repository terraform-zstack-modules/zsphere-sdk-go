// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateExternalPrimaryStorageParamDetail UpdateExternalPrimaryStorage detail param
type UpdateExternalPrimaryStorageParamDetail struct {
	Config *string `json:"config,omitempty"`
	DefaultProtocol *string `json:"defaultProtocol,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Url *string `json:"url,omitempty"`
}

// UpdateExternalPrimaryStorageParam UpdateExternalPrimaryStorage request param
type UpdateExternalPrimaryStorageParam struct {
	BaseParam
	Params UpdateExternalPrimaryStorageParamDetail `json:"updateExternalPrimaryStorage"`
}
// AddExternalPrimaryStorageParamDetail AddExternalPrimaryStorage detail param
type AddExternalPrimaryStorageParamDetail struct {
	Identity string `json:"identity" validate:"required"`
	DefaultOutputProtocol string `json:"defaultOutputProtocol" validate:"required"`
	Config *string `json:"config,omitempty"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddExternalPrimaryStorageParam AddExternalPrimaryStorage request param
type AddExternalPrimaryStorageParam struct {
	BaseParam
	Params AddExternalPrimaryStorageParamDetail `json:"params"`
}
