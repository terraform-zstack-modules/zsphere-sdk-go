// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VolumeTOView VolumeTO
type VolumeTOView struct {
	InstallPath string `json:"installPath,omitempty"`
	DeviceId int `json:"deviceId,omitempty"`
	DeviceType string `json:"deviceType,omitempty"`
	VolumeUuid string `json:"volumeUuid,omitempty"`
	UseVirtio bool `json:"useVirtio,omitempty"`
	UseVirtioSCSI bool `json:"useVirtioSCSI,omitempty"`
	UseSCSI bool `json:"useSCSI,omitempty"`
	Shareable bool `json:"shareable,omitempty"`
	CacheMode string `json:"cacheMode,omitempty"`
	AioNative bool `json:"aioNative,omitempty"`
	Wwn string `json:"wwn,omitempty"`
	BootOrder int `json:"bootOrder,omitempty"`
	PhysicalBlockSize int `json:"physicalBlockSize,omitempty"`
	Type string `json:"type,omitempty"`
	Format string `json:"format,omitempty"`
	PrimaryStorageType string `json:"primaryStorageType,omitempty"`
	MultiQueues string `json:"multiQueues,omitempty"`
	IoThreadId int `json:"ioThreadId,omitempty"`
	IoThreadPin string `json:"ioThreadPin,omitempty"`
	ControllerIndex int `json:"controllerIndex,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	DeviceAddress DeviceAddressView `json:"deviceAddress,omitempty"`
}

