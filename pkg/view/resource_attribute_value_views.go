// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ResourceAttributeValueInventoryView ResourceAttributeValue
type ResourceAttributeValueInventoryView struct {
	BaseInfoView
	BaseTimeView
	KeyUuid string `json:"keyUuid,omitempty"`
	Key ResourceAttributeKeyInventoryView `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
}

// DeleteResourceAttributeValueEventView DeleteResourceAttributeValueEvent
type DeleteResourceAttributeValueEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateResourceAttributeValueEventView CreateResourceAttributeValueEvent
type CreateResourceAttributeValueEventView struct {
	Inventories []CreateResourceAttributeResultView `json:"inventories,omitempty"`
}

// QueryResourceAttributeValueView QueryResourceAttributeValue
type QueryResourceAttributeValueView struct {
	Inventories []ResourceAttributeValueInventoryView `json:"inventories,omitempty"`
}

