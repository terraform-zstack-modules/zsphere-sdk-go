// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ResourceInventoryView Resource
type ResourceInventoryView struct {
	BaseInfoView
	BaseTimeView
	ResourceName string `json:"resourceName,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
}

// GetResourceNamesView GetResourceNames
type GetResourceNamesView struct {
	Inventories []ResourceInventoryView `json:"inventories,omitempty"`
}

