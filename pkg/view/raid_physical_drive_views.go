// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// RaidPhysicalDriveInventoryView RaidPhysicalDrive
type RaidPhysicalDriveInventoryView struct {
	BaseInfoView
	BaseTimeView
	RaidLevel string `json:"raidLevel,omitempty"`
	RaidControllerUuid string `json:"raidControllerUuid,omitempty"`
	Description string `json:"description,omitempty"`
	EnclosureDeviceId int `json:"enclosureDeviceId,omitempty"`
	SlotNumber int `json:"slotNumber,omitempty"`
	DeviceId int `json:"deviceId,omitempty"`
	DiskGroup int `json:"diskGroup,omitempty"`
	Wwn string `json:"wwn,omitempty"`
	SerialNumber string `json:"serialNumber,omitempty"`
	DeviceModel string `json:"deviceModel,omitempty"`
	Size int64 `json:"size,omitempty"`
	DriveState string `json:"driveState,omitempty"`
	LocateStatus string `json:"locateStatus,omitempty"`
	DriveType string `json:"driveType,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
	RotationRate int `json:"rotationRate,omitempty"`
}

// QueryLocalRaidPhysicalDriveView QueryLocalRaidPhysicalDrive
type QueryLocalRaidPhysicalDriveView struct {
	Inventories []RaidPhysicalDriveInventoryView `json:"inventories,omitempty"`
}

// LocateLocalRaidPhysicalDriveEventView LocateLocalRaidPhysicalDriveEvent
type LocateLocalRaidPhysicalDriveEventView struct {
	Inventory RaidPhysicalDriveInventoryView `json:"inventory,omitempty"`
}

