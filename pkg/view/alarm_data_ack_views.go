// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AlarmDataAckInventoryView AlarmDataAck
type AlarmDataAckInventoryView struct {
	BaseInfoView
	BaseTimeView
	AlarmUuid string `json:"alarmUuid,omitempty"`
	AlertDataUuid string `json:"alertDataUuid,omitempty"`
	AlertType string `json:"alertType,omitempty"`
	AckPeriod int64 `json:"ackPeriod,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	AckDate time.Time `json:"ackDate,omitempty"`
	ResumeAlert bool `json:"resumeAlert,omitempty"`
	OperatorAccountUuid string `json:"operatorAccountUuid,omitempty"`
}

