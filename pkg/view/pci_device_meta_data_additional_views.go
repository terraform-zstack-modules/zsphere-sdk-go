// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PciDeviceMetaDataView PciDeviceMetaData
type PciDeviceMetaDataView struct {
	MetaData string `json:"metaData,omitempty"`
	MetaDataEntries []PciDeviceMetaDataEntryView `json:"metaDataEntries,omitempty"`
}

