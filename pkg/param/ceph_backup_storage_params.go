// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// AddCephBackupStorageParamDetail AddCephBackupStorage detail param
type AddCephBackupStorageParamDetail struct {
	MonUrls []string `json:"monUrls" validate:"required"`
	PoolName *string `json:"poolName,omitempty"`
	Url *string `json:"url,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	ImportImages *bool `json:"importImages,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddCephBackupStorageParam AddCephBackupStorage request param
type AddCephBackupStorageParam struct {
	BaseParam
	Params AddCephBackupStorageParamDetail `json:"params"`
}
