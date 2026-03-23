// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AutoScalingGroupInventoryView AutoScalingGroup
type AutoScalingGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	ScalingResourceType string `json:"scalingResourceType,omitempty"`
	State string `json:"state,omitempty"`
	DefaultCooldown int64 `json:"defaultCooldown,omitempty"`
	Description string `json:"description,omitempty"`
	MinResourceSize int `json:"minResourceSize,omitempty"`
	MaxResourceSize int `json:"maxResourceSize,omitempty"`
	RemovalPolicy string `json:"removalPolicy,omitempty"`
	AttachedTemplates []string `json:"attachedTemplates,omitempty"`
	SystemTags []string `json:"systemTags,omitempty"`
}

// ChangeAutoScalingGroupStateEventView ChangeAutoScalingGroupStateEvent
type ChangeAutoScalingGroupStateEventView struct {
	Inventory AutoScalingGroupInventoryView `json:"inventory,omitempty"`
}

// UpdateAutoScalingGroupEventView UpdateAutoScalingGroupEvent
type UpdateAutoScalingGroupEventView struct {
	Inventory AutoScalingGroupInventoryView `json:"inventory,omitempty"`
}

// QueryAutoScalingGroupView QueryAutoScalingGroup
type QueryAutoScalingGroupView struct {
	Inventories []AutoScalingGroupInventoryView `json:"inventories,omitempty"`
}

// AttachAutoScalingTemplateToGroupEventView AttachAutoScalingTemplateToGroupEvent
type AttachAutoScalingTemplateToGroupEventView struct {
	Inventory AutoScalingGroupInventoryView `json:"inventory,omitempty"`
}

// DeleteAutoScalingGroupEventView DeleteAutoScalingGroupEvent
type DeleteAutoScalingGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// DetachAutoScalingTemplateFromGroupEventView DetachAutoScalingTemplateFromGroupEvent
type DetachAutoScalingTemplateFromGroupEventView struct {
	Inventory AutoScalingGroupInventoryView `json:"inventory,omitempty"`
}

// CreateAutoScalingGroupEventView CreateAutoScalingGroupEvent
type CreateAutoScalingGroupEventView struct {
	Inventory AutoScalingGroupInventoryView `json:"inventory,omitempty"`
}

