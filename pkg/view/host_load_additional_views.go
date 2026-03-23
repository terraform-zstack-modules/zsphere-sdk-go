// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostLoadView HostLoad
type HostLoadView struct {
	HostUuid string `json:"hostUuid,omitempty"`
	UsedCPUPercent float32 `json:"usedCPUPercent,omitempty"`
	UsedMemoryPercent float32 `json:"usedMemoryPercent,omitempty"`
}

