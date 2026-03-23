// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// EventLogInventoryView EventLog
type EventLogInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Category string `json:"category,omitempty"`
	TrackingId string `json:"trackingId,omitempty"`
	Type string `json:"type,omitempty"`
	Time int64 `json:"time,omitempty"`
}

// QueryEventLogView QueryEventLog
type QueryEventLogView struct {
	Inventories []EventLogInventoryView `json:"inventories,omitempty"`
}

