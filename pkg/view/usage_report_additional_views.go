// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// UsageReportView UsageReport
type UsageReportView struct {
	UsedPhysicalCapacitiesForecast []int64 `json:"usedPhysicalCapacitiesForecast,omitempty"`
	UsedPhysicalCapacitiesHistory []int64 `json:"usedPhysicalCapacitiesHistory,omitempty"`
	TotalPhysicalCapacitiesHistory []int64 `json:"totalPhysicalCapacitiesHistory,omitempty"`
	StartTime int64 `json:"startTime,omitempty"`
	Interval int64 `json:"interval,omitempty"`
}

