// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// OvfInfoView OvfInfo
type OvfInfoView struct {
	Disks []OvfDiskInfoView `json:"disks,omitempty"`
	Networks []OvfNetworkInfoView `json:"networks,omitempty"`
	Cpu OvfCpuInfoView `json:"cpu,omitempty"`
	Memory OvfMemoryInfoView `json:"memory,omitempty"`
	VmName string `json:"vmName,omitempty"`
	Os OvfOSInfoView `json:"os,omitempty"`
	SystemInfo OvfSystemInfoView `json:"systemInfo,omitempty"`
	Nics []OvfEthernetAdapterInfoView `json:"nics,omitempty"`
	CdDrivers []OvfCdDriverInfoView `json:"cdDrivers,omitempty"`
	Volumes []OvfVolumeInfoView `json:"volumes,omitempty"`
}

