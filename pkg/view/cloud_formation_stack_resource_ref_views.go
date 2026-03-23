// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CloudFormationStackResourceRefInventoryView CloudFormationStackResourceRef
type CloudFormationStackResourceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	StackUuid string `json:"stackUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Reserve bool `json:"reserve,omitempty"`
	Round int `json:"round,omitempty"`
}

