// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateZBoxBackupParamDetail CreateZBoxBackup detail param
type CreateZBoxBackupParamDetail struct {
	ZBoxUuid string `json:"zBoxUuid" validate:"required"`
	HostUuids []string `json:"hostUuids,omitempty"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
	DryRun *bool `json:"dryRun,omitempty"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateZBoxBackupParam CreateZBoxBackup request param
type CreateZBoxBackupParam struct {
	BaseParam
	Params CreateZBoxBackupParamDetail `json:"params"`
}
