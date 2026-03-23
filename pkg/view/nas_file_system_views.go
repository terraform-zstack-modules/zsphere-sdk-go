// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// NasFileSystemInventoryView NasFileSystem
type NasFileSystemInventoryView struct {
	BaseInfoView
	BaseTimeView
	Protocol string `json:"protocol,omitempty"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	FileSystemId string `json:"fileSystemId,omitempty"`
}

// QueryNasFileSystemView QueryNasFileSystem
type QueryNasFileSystemView struct {
	Inventories []NasFileSystemInventoryView `json:"inventories,omitempty"`
}

// UpdateNasFileSystemEventView UpdateNasFileSystemEvent
type UpdateNasFileSystemEventView struct {
	Inventory NasFileSystemInventoryView `json:"inventory,omitempty"`
}

// DeleteNasFileSystemEventView DeleteNasFileSystemEvent
type DeleteNasFileSystemEventView struct {
	Success bool `json:"success,omitempty"`
}

