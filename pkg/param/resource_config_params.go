// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// GetResourceConfigParamDetail GetResourceConfig detail param
type GetResourceConfigParamDetail struct {
}

// GetResourceConfigParam GetResourceConfig request param
type GetResourceConfigParam struct {
	BaseParam
	Params GetResourceConfigParamDetail `json:"getResourceConfig"`
}
// UpdateResourceConfigParamDetail UpdateResourceConfig detail param
type UpdateResourceConfigParamDetail struct {
	Value string `json:"value" validate:"required"`
}

// UpdateResourceConfigParam UpdateResourceConfig request param
type UpdateResourceConfigParam struct {
	BaseParam
	Params UpdateResourceConfigParamDetail `json:"updateResourceConfig"`
}
// DeleteResourceConfigParamDetail DeleteResourceConfig detail param
type DeleteResourceConfigParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteResourceConfigParam DeleteResourceConfig request param
type DeleteResourceConfigParam struct {
	BaseParam
	Params DeleteResourceConfigParamDetail `json:"deleteResourceConfig"`
}
