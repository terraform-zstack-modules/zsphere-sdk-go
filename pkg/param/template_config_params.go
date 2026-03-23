// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateTemplateConfigParamDetail UpdateTemplateConfig detail param
type UpdateTemplateConfigParamDetail struct {
	Category string `json:"category" validate:"required"`
	Name string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// UpdateTemplateConfigParam UpdateTemplateConfig request param
type UpdateTemplateConfigParam struct {
	BaseParam
	Params UpdateTemplateConfigParamDetail `json:"updateTemplateConfig"`
}
// ApplyTemplateConfigParamDetail ApplyTemplateConfig detail param
type ApplyTemplateConfigParamDetail struct {
}

// ApplyTemplateConfigParam ApplyTemplateConfig request param
type ApplyTemplateConfigParam struct {
	BaseParam
	Params ApplyTemplateConfigParamDetail `json:"applyTemplateConfig"`
}
// RevertTemplateConfigParamDetail RevertTemplateConfig detail param
type RevertTemplateConfigParamDetail struct {
}

// RevertTemplateConfigParam RevertTemplateConfig request param
type RevertTemplateConfigParam struct {
	BaseParam
	Params RevertTemplateConfigParamDetail `json:"revertTemplateConfig"`
}
// ResetTemplateConfigParamDetail ResetTemplateConfig detail param
type ResetTemplateConfigParamDetail struct {
}

// ResetTemplateConfigParam ResetTemplateConfig request param
type ResetTemplateConfigParam struct {
	BaseParam
	Params ResetTemplateConfigParamDetail `json:"resetTemplateConfig"`
}
