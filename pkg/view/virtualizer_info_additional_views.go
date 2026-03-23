// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VirtualizerInfoView VirtualizerInfo
type VirtualizerInfoView struct {
	Hypervisor string `json:"hypervisor,omitempty"`
	CurrentVersion string `json:"currentVersion,omitempty"`
	ExpectVersion string `json:"expectVersion,omitempty"`
	MatchState string `json:"matchState,omitempty"`
}

// VirtualizerInfoInventoryView VirtualizerInfo
type VirtualizerInfoInventoryView struct {
	BaseInfoView
	BaseTimeView
	ResourceType string `json:"resourceType,omitempty"`
	InfoList []VirtualizerInfoView `json:"infoList,omitempty"`
}

