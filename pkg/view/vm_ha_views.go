// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmHaInventoryView VmHa
type VmHaInventoryView struct {
	BaseInfoView
	BaseTimeView
	HaLevel string `json:"haLevel,omitempty"`
	HaLevelUpdateTime time.Time `json:"haLevelUpdateTime,omitempty"`
	InhibitionReason string `json:"inhibitionReason,omitempty"`
	InhibitionTime time.Time `json:"inhibitionTime,omitempty"`
}

