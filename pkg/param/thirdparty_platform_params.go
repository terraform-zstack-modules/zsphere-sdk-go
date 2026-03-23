// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateThirdpartyPlatformParamDetail UpdateThirdpartyPlatform detail param
type UpdateThirdpartyPlatformParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Template *string `json:"template,omitempty"`
	Url *string `json:"url,omitempty"`
	StateEvent *string `json:"stateEvent,omitempty"`
	LastSyncDateMills *int64 `json:"lastSyncDateMills,omitempty"`
}

// UpdateThirdpartyPlatformParam UpdateThirdpartyPlatform request param
type UpdateThirdpartyPlatformParam struct {
	BaseParam
	Params UpdateThirdpartyPlatformParamDetail `json:"updateThirdpartyPlatform"`
}
// AddThirdpartyPlatformParamDetail AddThirdpartyPlatform detail param
type AddThirdpartyPlatformParamDetail struct {
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required"`
	Url string `json:"url" validate:"required"`
	Template string `json:"template" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddThirdpartyPlatformParam AddThirdpartyPlatform request param
type AddThirdpartyPlatformParam struct {
	BaseParam
	Params AddThirdpartyPlatformParamDetail `json:"params"`
}
// DeleteThirdpartyPlatformParamDetail DeleteThirdpartyPlatform detail param
type DeleteThirdpartyPlatformParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteThirdpartyPlatformParam DeleteThirdpartyPlatform request param
type DeleteThirdpartyPlatformParam struct {
	BaseParam
	Params DeleteThirdpartyPlatformParamDetail `json:"deleteThirdpartyPlatform"`
}
