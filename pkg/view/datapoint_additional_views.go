// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// DatapointView Datapoint
type DatapointView struct {
	Value float64 `json:"value,omitempty"`
	Time int64 `json:"time,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
}

