// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VolumeCbtBackupInfoView VolumeCbtBackupInfo
type VolumeCbtBackupInfoView struct {
	Volume VolumeTOView `json:"volume,omitempty"`
	BitmapBase64 string `json:"bitmapBase64,omitempty"`
	Target string `json:"target,omitempty"`
	ScratchNodeName string `json:"scratchNodeName,omitempty"`
	Metadata string `json:"metadata,omitempty"`
	NbdPort int64 `json:"nbdPort,omitempty"`
	NbdServer string `json:"nbdServer,omitempty"`
	Mode string `json:"mode,omitempty"`
	BitmapName string `json:"bitmapName,omitempty"`
}

