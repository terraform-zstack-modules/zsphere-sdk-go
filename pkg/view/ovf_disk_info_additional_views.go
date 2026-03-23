// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// OvfDiskInfoView OvfDiskInfo
type OvfDiskInfoView struct {
	Index int `json:"index,omitempty"`
	DiskId string `json:"diskId,omitempty"`
	FileRef string `json:"fileRef,omitempty"`
	FileName string `json:"fileName,omitempty"`
	Format string `json:"format,omitempty"`
	PopulatedSize int64 `json:"populatedSize,omitempty"`
	Capacity int64 `json:"capacity,omitempty"`
}

