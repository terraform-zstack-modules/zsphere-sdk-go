// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AffinityGroupUsageInventoryView AffinityGroupUsage
type AffinityGroupUsageInventoryView struct {
	BaseInfoView
	BaseTimeView
	AffinityGroupUuid string `json:"affinityGroupUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
}

