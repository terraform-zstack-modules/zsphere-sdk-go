// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteAutoScalingTemplateParamDetail DeleteAutoScalingTemplate detail param
type DeleteAutoScalingTemplateParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAutoScalingTemplateParam DeleteAutoScalingTemplate request param
type DeleteAutoScalingTemplateParam struct {
	BaseParam
	Params DeleteAutoScalingTemplateParamDetail `json:"deleteAutoScalingTemplate"`
}
