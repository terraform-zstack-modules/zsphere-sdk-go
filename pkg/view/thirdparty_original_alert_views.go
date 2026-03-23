// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ThirdpartyOriginalAlertInventoryView ThirdpartyOriginalAlert
type ThirdpartyOriginalAlertInventoryView struct {
	BaseInfoView
	BaseTimeView
	ThirdpartyPlatformUuid string `json:"thirdpartyPlatformUuid,omitempty"`
	Product string `json:"product,omitempty"`
	Service string `json:"service,omitempty"`
	Metric string `json:"metric,omitempty"`
	AlertLevel string `json:"alertLevel,omitempty"`
	AlertTime time.Time `json:"alertTime,omitempty"`
	Dimensions string `json:"dimensions,omitempty"`
	Message string `json:"message,omitempty"`
	DataSource string `json:"dataSource,omitempty"`
	SourceText string `json:"sourceText,omitempty"`
	ReadStatus string `json:"readStatus,omitempty"`
}

// QueryThirdpartyAlertView QueryThirdpartyAlert
type QueryThirdpartyAlertView struct {
	Inventories []ThirdpartyOriginalAlertInventoryView `json:"inventories,omitempty"`
}

