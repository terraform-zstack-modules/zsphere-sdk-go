// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HistogramView Histogram
type HistogramView struct {
	Time int64 `json:"time,omitempty"`
	Count int64 `json:"count,omitempty"`
	Tags []TagView `json:"tags,omitempty"`
}

