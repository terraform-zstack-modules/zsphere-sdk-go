// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BlockDeviceView BlockDevice
type BlockDeviceView struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Size int64 `json:"size,omitempty"`
	Used int64 `json:"used,omitempty"`
	Available int64 `json:"available,omitempty"`
	PhysicalSector int64 `json:"physicalSector,omitempty"`
	LogicalSector int64 `json:"logicalSector,omitempty"`
	MountPoint string `json:"mountPoint,omitempty"`
	Children []*BlockDeviceView `json:"children,omitempty"`
	PartitionTable string `json:"partitionTable,omitempty"`
	FsType string `json:"fsType,omitempty"`
	SerialNumber string `json:"serialNumber,omitempty"`
	Model string `json:"model,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
	UsedRatio int64 `json:"usedRatio,omitempty"`
	SmartPassed bool `json:"smartPassed,omitempty"`
	SmartMessage string `json:"smartMessage,omitempty"`
}

