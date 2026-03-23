// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LongJobInventoryView LongJob
type LongJobInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ApiId string `json:"apiId,omitempty"`
	JobName string `json:"jobName,omitempty"`
	JobData string `json:"jobData,omitempty"`
	JobResult string `json:"jobResult,omitempty"`
	State string `json:"state,omitempty"`
	TargetResourceUuid string `json:"targetResourceUuid,omitempty"`
	ManagementNodeUuid string `json:"managementNodeUuid,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	ExecuteTime int64 `json:"executeTime,omitempty"`
}

// CleanLongJobEventView CleanLongJobEvent
type CleanLongJobEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateClusterOSEventView UpdateClusterOSEvent
type UpdateClusterOSEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// ResumeLongJobEventView ResumeLongJobEvent
type ResumeLongJobEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// BatchCreateBaremetalChassisEventView BatchCreateBaremetalChassisEvent
type BatchCreateBaremetalChassisEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// DeleteLongJobEventView DeleteLongJobEvent
type DeleteLongJobEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateLongJobEventView UpdateLongJobEvent
type UpdateLongJobEventView struct {
	Inventory LongJobInventoryView `json:"inventory,omitempty"`
}

// QueryLongJobView QueryLongJob
type QueryLongJobView struct {
	Inventories []LongJobInventoryView `json:"inventories,omitempty"`
}

