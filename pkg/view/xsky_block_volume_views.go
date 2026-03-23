// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// XskyBlockVolumeInventoryView XskyBlockVolume
type XskyBlockVolumeInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccessPathId int `json:"accessPathId,omitempty"`
	AccessPathIqn string `json:"accessPathIqn,omitempty"`
	XskyStatus string `json:"xskyStatus,omitempty"`
	XskyBlockVolumeId int `json:"xskyBlockVolumeId,omitempty"`
	BurstTotalBw int64 `json:"burstTotalBw,omitempty"`
	BurstTotalIops int64 `json:"burstTotalIops,omitempty"`
	MaxTotalBw int64 `json:"maxTotalBw,omitempty"`
	MaxTotalIops int64 `json:"maxTotalIops,omitempty"`
	IscsiPath string `json:"iscsiPath,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	Description string `json:"description,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	DiskOfferingUuid string `json:"diskOfferingUuid,omitempty"`
	RootImageUuid string `json:"rootImageUuid,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	Type string `json:"type,omitempty"`
	Format string `json:"format,omitempty"`
	Size int64 `json:"size,omitempty"`
	ActualSize int64 `json:"actualSize,omitempty"`
	DeviceId int `json:"deviceId,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	IsShareable bool `json:"isShareable,omitempty"`
	VolumeQos string `json:"volumeQos,omitempty"`
	LastDetachDate time.Time `json:"lastDetachDate,omitempty"`
	LastVmInstanceUuid string `json:"lastVmInstanceUuid,omitempty"`
	LastAttachDate time.Time `json:"lastAttachDate,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

// QueryXskyBlockVolumeView QueryXskyBlockVolume
type QueryXskyBlockVolumeView struct {
	Inventories []XskyBlockVolumeInventoryView `json:"inventories,omitempty"`
}

