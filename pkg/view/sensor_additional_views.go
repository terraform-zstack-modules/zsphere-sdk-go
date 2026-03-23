// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SensorView Sensor
type SensorView struct {
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Status string `json:"status,omitempty"`
	Type string `json:"type,omitempty"`
	Classification string `json:"classification,omitempty"`
}

