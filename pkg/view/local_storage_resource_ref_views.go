// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LocalStorageResourceRefInventoryView LocalStorageResourceRef
type LocalStorageResourceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	ResourceUuid string `json:"resourceUuid,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	Size int64 `json:"size,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
}

// QueryLocalStorageResourceRefView QueryLocalStorageResourceRef
type QueryLocalStorageResourceRefView struct {
	Inventories []LocalStorageResourceRefInventoryView `json:"inventories,omitempty"`
}

