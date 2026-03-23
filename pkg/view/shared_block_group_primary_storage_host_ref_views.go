// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SharedBlockGroupPrimaryStorageHostRefInventoryView SharedBlockGroupPrimaryStorageHostRef
type SharedBlockGroupPrimaryStorageHostRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	HostId int `json:"hostId,omitempty"`
	Status string `json:"status,omitempty"`
}

// QuerySharedBlockGroupPrimaryStorageHostRefView QuerySharedBlockGroupPrimaryStorageHostRef
type QuerySharedBlockGroupPrimaryStorageHostRefView struct {
	Inventories []SharedBlockGroupPrimaryStorageHostRefInventoryView `json:"inventories,omitempty"`
}

