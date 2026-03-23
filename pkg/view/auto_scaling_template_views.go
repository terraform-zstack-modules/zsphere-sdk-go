// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AutoScalingTemplateInventoryView AutoScalingTemplate
type AutoScalingTemplateInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	SystemTags []string `json:"systemTags,omitempty"`
}

// DeleteAutoScalingTemplateEventView DeleteAutoScalingTemplateEvent
type DeleteAutoScalingTemplateEventView struct {
	Success bool `json:"success,omitempty"`
}

