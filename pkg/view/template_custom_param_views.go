// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// TemplateCustomParamInventoryView TemplateCustomParam
type TemplateCustomParamInventoryView struct {
	BaseInfoView
	BaseTimeView
	TemplateUuid string `json:"templateUuid,omitempty"`
	Param string `json:"param,omitempty"`
}

