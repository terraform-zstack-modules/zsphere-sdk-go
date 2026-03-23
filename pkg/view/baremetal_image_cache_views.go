// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BaremetalImageCacheInventoryView BaremetalImageCache
type BaremetalImageCacheInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	PxeServerUuid string `json:"pxeServerUuid,omitempty"`
	ImageUuid string `json:"imageUuid,omitempty"`
	Url string `json:"url,omitempty"`
	InstallUrl string `json:"installUrl,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
	Size int64 `json:"size,omitempty"`
	ActualSize int64 `json:"actualSize,omitempty"`
	Md5sum string `json:"md5sum,omitempty"`
	Utilization int64 `json:"utilization,omitempty"`
}

