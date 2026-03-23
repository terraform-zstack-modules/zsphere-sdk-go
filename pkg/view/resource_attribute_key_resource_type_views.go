// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ResourceAttributeKeyResourceTypeInventoryView ResourceAttributeKeyResourceType
type ResourceAttributeKeyResourceTypeInventoryView struct {
	BaseInfoView
	BaseTimeView
	KeyUuid string `json:"keyUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
}

