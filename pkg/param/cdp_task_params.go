// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateCdpTaskParamDetail CreateCdpTask detail param
type CreateCdpTaskParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	TaskType string `json:"taskType" validate:"required"`
	PolicyUuid string `json:"policyUuid" validate:"required"`
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
	BackupBandwidth *int64 `json:"backupBandwidth,omitempty"`
	MaxCapacity *int64 `json:"maxCapacity,omitempty"`
	MaxLatency *int64 `json:"maxLatency,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateCdpTaskParam CreateCdpTask request param
type CreateCdpTaskParam struct {
	BaseParam
	Params CreateCdpTaskParamDetail `json:"params"`
}
// DeleteCdpTaskParamDetail DeleteCdpTask detail param
type DeleteCdpTaskParamDetail struct {
	Force *bool `json:"force,omitempty"`
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteCdpTaskParam DeleteCdpTask request param
type DeleteCdpTaskParam struct {
	BaseParam
	Params DeleteCdpTaskParamDetail `json:"deleteCdpTask"`
}
// UpdateCdpTaskParamDetail UpdateCdpTask detail param
type UpdateCdpTaskParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	BackupBandwidth *int64 `json:"backupBandwidth,omitempty"`
	MaxCapacity *int64 `json:"maxCapacity,omitempty"`
	MaxLatency *int64 `json:"maxLatency,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// UpdateCdpTaskParam UpdateCdpTask request param
type UpdateCdpTaskParam struct {
	BaseParam
	Params UpdateCdpTaskParamDetail `json:"updateCdpTask"`
}
