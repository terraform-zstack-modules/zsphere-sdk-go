// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// DatabaseBackupStructView DatabaseBackupStruct
type DatabaseBackupStructView struct {
	Name string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	Type string `json:"type,omitempty"`
	Size int64 `json:"size,omitempty"`
	Md5 string `json:"md5,omitempty"`
	CreatedTime time.Time `json:"createdTime,omitempty"`
}

