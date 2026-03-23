// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VolumeCbtBackupRecordInventoryView VolumeCbtBackupRecord
type VolumeCbtBackupRecordInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	TaskUuid string `json:"taskUuid,omitempty"`
	VolumeUuid string `json:"volumeUuid,omitempty"`
	Mode string `json:"mode,omitempty"`
	Target string `json:"target,omitempty"`
	ScratchNodeName string `json:"scratchNodeName,omitempty"`
	BitmapName string `json:"bitmapName,omitempty"`
	LastBitmapName string `json:"lastBitmapName,omitempty"`
}

