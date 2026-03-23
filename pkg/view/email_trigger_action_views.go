// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// EmailTriggerActionInventoryView EmailTriggerAction
type EmailTriggerActionInventoryView struct {
	BaseInfoView
	BaseTimeView
	Email string `json:"email,omitempty"`
	MediaUuid string `json:"mediaUuid,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
	TriggerUuids []string `json:"triggerUuids,omitempty"`
}

