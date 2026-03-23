// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostOsCategoryInventoryView HostOsCategory
type HostOsCategoryInventoryView struct {
	BaseInfoView
	BaseTimeView
	Architecture string `json:"architecture,omitempty"`
	OsReleaseVersion string `json:"osReleaseVersion,omitempty"`
	MetadataList []KvmHostHypervisorMetadataInventoryView `json:"metadataList,omitempty"`
}

// QueryHostOsCategoryView QueryHostOsCategory
type QueryHostOsCategoryView struct {
	Inventories []HostOsCategoryInventoryView `json:"inventories,omitempty"`
}

