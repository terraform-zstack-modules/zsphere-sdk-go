// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// AddCephPrimaryStorageParamDetail AddCephPrimaryStorage detail param
type AddCephPrimaryStorageParamDetail struct {
	MonUrls []string `json:"monUrls" validate:"required"`
	RootVolumePoolName *string `json:"rootVolumePoolName,omitempty"`
	DataVolumePoolName *string `json:"dataVolumePoolName,omitempty"`
	ImageCachePoolName *string `json:"imageCachePoolName,omitempty"`
	Url *string `json:"url,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddCephPrimaryStorageParam AddCephPrimaryStorage request param
type AddCephPrimaryStorageParam struct {
	BaseParam
	Params AddCephPrimaryStorageParamDetail `json:"params"`
}
