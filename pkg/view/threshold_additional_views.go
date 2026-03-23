// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ThresholdView Threshold
type ThresholdView struct {
	ThresholdName string `json:"thresholdName,omitempty"`
	ThresholdValue string `json:"thresholdValue,omitempty"`
	Operator string `json:"operator,omitempty"`
}

