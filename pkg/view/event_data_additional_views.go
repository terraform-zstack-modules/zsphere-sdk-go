// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// EventDataView EventData
type EventDataView struct {
	Namespace string `json:"namespace,omitempty"`
	Name string `json:"name,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	ResourceId string `json:"resourceId,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
	Time int64 `json:"time,omitempty"`
	DataUuid string `json:"dataUuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	SubscriptionUuid string `json:"subscriptionUuid,omitempty"`
	ReadStatus string `json:"readStatus,omitempty"`
}

