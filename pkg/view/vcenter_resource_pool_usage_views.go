// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VCenterResourcePoolUsageInventoryView VCenterResourcePoolUsage
type VCenterResourcePoolUsageInventoryView struct {
	BaseInfoView
	BaseTimeView
	VCenterResourcePoolUuid string `json:"vCenterResourcePoolUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
}

