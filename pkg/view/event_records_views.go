// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// EventRecordsInventoryView EventRecords
type EventRecordsInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	CreateTime int64 `json:"createTime,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	ResourceId string `json:"resourceId,omitempty"`
	DataUuid string `json:"dataUuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	SubscriptionUuid string `json:"subscriptionUuid,omitempty"`
	ReadStatus bool `json:"readStatus,omitempty"`
	OperatorAccountUuid string `json:"operatorAccountUuid,omitempty"`
	Labels string `json:"labels,omitempty"`
}

// QueryEventRecordView QueryEventRecord
type QueryEventRecordView struct {
	Inventories []EventRecordsInventoryView `json:"inventories,omitempty"`
}

