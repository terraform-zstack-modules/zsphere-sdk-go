// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteVolumeSnapshotParamDetail DeleteVolumeSnapshot detail param
type DeleteVolumeSnapshotParamDetail struct {
	Direction *string `json:"direction,omitempty"`
	Scope *string `json:"scope,omitempty"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVolumeSnapshotParam DeleteVolumeSnapshot request param
type DeleteVolumeSnapshotParam struct {
	BaseParam
	Params DeleteVolumeSnapshotParamDetail `json:"deleteVolumeSnapshot"`
}
// CreateVolumeSnapshotParamDetail CreateVolumeSnapshot detail param
type CreateVolumeSnapshotParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVolumeSnapshotParam CreateVolumeSnapshot request param
type CreateVolumeSnapshotParam struct {
	BaseParam
	Params CreateVolumeSnapshotParamDetail `json:"params"`
}
// UpdateVolumeSnapshotParamDetail UpdateVolumeSnapshot detail param
type UpdateVolumeSnapshotParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateVolumeSnapshotParam UpdateVolumeSnapshot request param
type UpdateVolumeSnapshotParam struct {
	BaseParam
	Params UpdateVolumeSnapshotParamDetail `json:"updateVolumeSnapshot"`
}
