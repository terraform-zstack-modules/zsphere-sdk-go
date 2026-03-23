// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AutoScalingTemplateGroupRefInventoryView AutoScalingTemplateGroupRef
type AutoScalingTemplateGroupRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	TemplateUuid string `json:"templateUuid,omitempty"`
	GroupUuid string `json:"groupUuid,omitempty"`
}

