// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AlertDataAckInventoryView AlertDataAck
type AlertDataAckInventoryView struct {
	BaseInfoView
	BaseTimeView
	AlertDataUuid string `json:"alertDataUuid,omitempty"`
	AlertType string `json:"alertType,omitempty"`
	AckPeriod int64 `json:"ackPeriod,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	AckDate time.Time `json:"ackDate,omitempty"`
	ResumeAlert bool `json:"resumeAlert,omitempty"`
	OperatorAccountUuid string `json:"operatorAccountUuid,omitempty"`
}

// UpdateAlertDataAckEventView UpdateAlertDataAckEvent
type UpdateAlertDataAckEventView struct {
	Inventory AlertDataAckInventoryView `json:"inventory,omitempty"`
}

// QueryAlertDataAckView QueryAlertDataAck
type QueryAlertDataAckView struct {
	Inventories []AlertDataAckInventoryView `json:"inventories,omitempty"`
}

// AckAlertDataEventView AckAlertDataEvent
type AckAlertDataEventView struct {
	Inventory AlertDataAckInventoryView `json:"inventory,omitempty"`
}

