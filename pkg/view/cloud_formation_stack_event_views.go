// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CloudFormationStackEventInventoryView CloudFormationStackEvent
type CloudFormationStackEventInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Action string `json:"action,omitempty"`
	Content string `json:"content,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
	ActionStatus string `json:"actionStatus,omitempty"`
	StackUuid string `json:"stackUuid,omitempty"`
	Duration string `json:"duration,omitempty"`
}

// QueryEventFromResourceStackView QueryEventFromResourceStack
type QueryEventFromResourceStackView struct {
	Inventories []CloudFormationStackEventInventoryView `json:"inventories,omitempty"`
}

