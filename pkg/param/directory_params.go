// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateDirectoryParamDetail UpdateDirectory detail param
type UpdateDirectoryParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name" validate:"required"`
}

// UpdateDirectoryParam UpdateDirectory request param
type UpdateDirectoryParam struct {
	BaseParam
	Params UpdateDirectoryParamDetail `json:"updateDirectory"`
}
// DeleteDirectoryParamDetail DeleteDirectory detail param
type DeleteDirectoryParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteDirectoryParam DeleteDirectory request param
type DeleteDirectoryParam struct {
	BaseParam
	Params DeleteDirectoryParamDetail `json:"deleteDirectory"`
}
// CreateDirectoryParamDetail CreateDirectory detail param
type CreateDirectoryParamDetail struct {
	Name string `json:"name" validate:"required"`
	ParentUuid *string `json:"parentUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	Type string `json:"type" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDirectoryParam CreateDirectory request param
type CreateDirectoryParam struct {
	BaseParam
	Params CreateDirectoryParamDetail `json:"params"`
}
