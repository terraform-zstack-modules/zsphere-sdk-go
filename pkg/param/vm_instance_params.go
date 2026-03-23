// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CloneVmInstanceParamDetail CloneVmInstance detail param
type CloneVmInstanceParamDetail struct {
	Strategy *string `json:"strategy,omitempty"`
	VmNicParams *string `json:"vmNicParams,omitempty"`
	Names []string `json:"names" validate:"required"`
	PrimaryStorageUuidForRootVolume *string `json:"primaryStorageUuidForRootVolume,omitempty"`
	PrimaryStorageUuidForDataVolume *string `json:"primaryStorageUuidForDataVolume,omitempty"`
	Full *bool `json:"full,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags []string `json:"dataVolumeSystemTags,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	DiskAOs []DiskAOParam `json:"diskAOs,omitempty"`
	Description *string `json:"description,omitempty"`
	VmCustomSpecification VmCustomSpecificationStructParam `json:"vmCustomSpecification,omitempty"`
}

// CloneVmInstanceParam CloneVmInstance request param
type CloneVmInstanceParam struct {
	BaseParam
	Params CloneVmInstanceParamDetail `json:"cloneVmInstance"`
}
// StopVmInstanceParamDetail StopVmInstance detail param
type StopVmInstanceParamDetail struct {
	Type *string `json:"type,omitempty"`
	StopHA *string `json:"stopHA,omitempty"`
}

// StopVmInstanceParam StopVmInstance request param
type StopVmInstanceParam struct {
	BaseParam
	Params StopVmInstanceParamDetail `json:"stopVmInstance"`
}
// ResumeVmInstanceParamDetail ResumeVmInstance detail param
type ResumeVmInstanceParamDetail struct {
}

// ResumeVmInstanceParam ResumeVmInstance request param
type ResumeVmInstanceParam struct {
	BaseParam
	Params ResumeVmInstanceParamDetail `json:"resumeVmInstance"`
}
// ExpungeVmInstanceParamDetail ExpungeVmInstance detail param
type ExpungeVmInstanceParamDetail struct {
}

// ExpungeVmInstanceParam ExpungeVmInstance request param
type ExpungeVmInstanceParam struct {
	BaseParam
	Params ExpungeVmInstanceParamDetail `json:"expungeVmInstance"`
}
// RebootVmInstanceParamDetail RebootVmInstance detail param
type RebootVmInstanceParamDetail struct {
}

// RebootVmInstanceParam RebootVmInstance request param
type RebootVmInstanceParam struct {
	BaseParam
	Params RebootVmInstanceParamDetail `json:"rebootVmInstance"`
}
// UpdateVmInstanceParamDetail UpdateVmInstance detail param
type UpdateVmInstanceParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	State *string `json:"state,omitempty"`
	DefaultL3NetworkUuid *string `json:"defaultL3NetworkUuid,omitempty"`
	Platform *string `json:"platform,omitempty"`
	CpuNum *int `json:"cpuNum,omitempty"`
	MemorySize *int64 `json:"memorySize,omitempty"`
	ReservedMemorySize *int64 `json:"reservedMemorySize,omitempty"`
	GuestOsType *string `json:"guestOsType,omitempty"`
	AllocatorStrategy *string `json:"allocatorStrategy,omitempty"`
}

// UpdateVmInstanceParam UpdateVmInstance request param
type UpdateVmInstanceParam struct {
	BaseParam
	Params UpdateVmInstanceParamDetail `json:"updateVmInstance"`
}
// DestroyVmInstanceParamDetail DestroyVmInstance detail param
type DestroyVmInstanceParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DestroyVmInstanceParam DestroyVmInstance request param
type DestroyVmInstanceParam struct {
	BaseParam
	Params DestroyVmInstanceParamDetail `json:"destroyVmInstance"`
}
// StartVmInstanceParamDetail StartVmInstance detail param
type StartVmInstanceParamDetail struct {
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
}

// StartVmInstanceParam StartVmInstance request param
type StartVmInstanceParam struct {
	BaseParam
	Params StartVmInstanceParamDetail `json:"startVmInstance"`
}
// CreateVmInstanceParamDetail CreateVmInstance detail param
type CreateVmInstanceParamDetail struct {
	Name string `json:"name" validate:"required"`
	InstanceOfferingUuid *string `json:"instanceOfferingUuid,omitempty"`
	CpuNum *int `json:"cpuNum,omitempty"`
	MemorySize *int64 `json:"memorySize,omitempty"`
	ReservedMemorySize *int64 `json:"reservedMemorySize,omitempty"`
	ImageUuid *string `json:"imageUuid,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	VmNicParams *string `json:"vmNicParams,omitempty"`
	Type *string `json:"type,omitempty"`
	RootDiskOfferingUuid *string `json:"rootDiskOfferingUuid,omitempty"`
	RootDiskSize *int64 `json:"rootDiskSize,omitempty"`
	DataDiskSizes []int64 `json:"dataDiskSizes,omitempty"`
	DataDiskOfferingUuids []string `json:"dataDiskOfferingUuids,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	HostUuid *string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume *string `json:"primaryStorageUuidForRootVolume,omitempty"`
	Description *string `json:"description,omitempty"`
	DefaultL3NetworkUuid *string `json:"defaultL3NetworkUuid,omitempty"`
	Strategy *string `json:"strategy,omitempty"`
	RootVolumeSystemTags []string `json:"rootVolumeSystemTags,omitempty"`
	DataVolumeSystemTags []string `json:"dataVolumeSystemTags,omitempty"`
	DataVolumeSystemTagsOnIndex map[string]interface{} `json:"dataVolumeSystemTagsOnIndex,omitempty"`
	SshKeyPairUuids []string `json:"sshKeyPairUuids,omitempty"`
	Platform *string `json:"platform,omitempty"`
	GuestOsType *string `json:"guestOsType,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	Virtio *bool `json:"virtio,omitempty"`
	AllocatorStrategy *string `json:"allocatorStrategy,omitempty"`
	DiskAOs []DiskAOParam `json:"diskAOs,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmInstanceParam CreateVmInstance request param
type CreateVmInstanceParam struct {
	BaseParam
	Params CreateVmInstanceParamDetail `json:"params"`
}
// RecoverVmInstanceParamDetail RecoverVmInstance detail param
type RecoverVmInstanceParamDetail struct {
}

// RecoverVmInstanceParam RecoverVmInstance request param
type RecoverVmInstanceParam struct {
	BaseParam
	Params RecoverVmInstanceParamDetail `json:"recoverVmInstance"`
}
