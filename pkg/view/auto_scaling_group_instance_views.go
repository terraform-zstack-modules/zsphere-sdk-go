// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AutoScalingGroupInstanceInventoryView AutoScalingGroupInstance
type AutoScalingGroupInstanceInventoryView struct {
	BaseInfoView
	BaseTimeView
	InstanceUuid string `json:"instanceUuid,omitempty"`
	ScalingGroupUuid string `json:"scalingGroupUuid,omitempty"`
	TemplateUuid string `json:"templateUuid,omitempty"`
	ScalingGroupActivityUuid string `json:"scalingGroupActivityUuid,omitempty"`
	Status string `json:"status,omitempty"`
	HealthStatus string `json:"healthStatus,omitempty"`
	Description string `json:"description,omitempty"`
	ProtectionStrategy string `json:"protectionStrategy,omitempty"`
}

// UpdateAutoScalingGroupInstanceEventView UpdateAutoScalingGroupInstanceEvent
type UpdateAutoScalingGroupInstanceEventView struct {
	Inventory AutoScalingGroupInstanceInventoryView `json:"inventory,omitempty"`
}

// QueryAutoScalingGroupInstanceView QueryAutoScalingGroupInstance
type QueryAutoScalingGroupInstanceView struct {
	Inventories []AutoScalingGroupInstanceInventoryView `json:"inventories,omitempty"`
}

// DeleteAutoScalingGroupInstanceEventView DeleteAutoScalingGroupInstanceEvent
type DeleteAutoScalingGroupInstanceEventView struct {
	Success bool `json:"success,omitempty"`
}

