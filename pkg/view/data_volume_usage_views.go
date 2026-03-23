// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// DataVolumeUsageInventoryView DataVolumeUsage
type DataVolumeUsageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	DateInLong int64 `json:"dateInLong,omitempty"`
	VolumeUuid string `json:"volumeUuid,omitempty"`
	VolumeStatus string `json:"volumeStatus,omitempty"`
	VolumeName string `json:"volumeName,omitempty"`
	VolumeSize int64 `json:"volumeSize,omitempty"`
	Inventory string `json:"inventory,omitempty"`
}

