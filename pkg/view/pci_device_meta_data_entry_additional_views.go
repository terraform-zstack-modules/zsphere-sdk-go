// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PciDeviceMetaDataEntryView PciDeviceMetaDataEntry
type PciDeviceMetaDataEntryView struct {
	Key string `json:"key,omitempty"`
	Op string `json:"op,omitempty"`
	Value string `json:"value,omitempty"`
}

