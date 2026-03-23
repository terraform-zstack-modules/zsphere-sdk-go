// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// RootVolumeBillingInventoryView RootVolumeBilling
type RootVolumeBillingInventoryView struct {
	BaseInfoView
	BaseTimeView
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	VolumeSize int64 `json:"volumeSize,omitempty"`
	Id int64 `json:"id,omitempty"`
	BillingType string `json:"billingType,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
	Spending float64 `json:"spending,omitempty"`
	StartTime int64 `json:"startTime,omitempty"`
	EndTime int64 `json:"endTime,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
}

