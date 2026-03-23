// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CdpTaskResourceRefInventoryView CdpTaskResourceRef
type CdpTaskResourceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	TaskUuid string `json:"taskUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
}

