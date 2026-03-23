// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateBackupStorageParamDetail UpdateBackupStorage detail param
type UpdateBackupStorageParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateBackupStorageParam UpdateBackupStorage request param
type UpdateBackupStorageParam struct {
	BaseParam
	Params UpdateBackupStorageParamDetail `json:"updateBackupStorage"`
}
// DeleteBackupStorageParamDetail DeleteBackupStorage detail param
type DeleteBackupStorageParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteBackupStorageParam DeleteBackupStorage request param
type DeleteBackupStorageParam struct {
	BaseParam
	Params DeleteBackupStorageParamDetail `json:"deleteBackupStorage"`
}
// ReconnectBackupStorageParamDetail ReconnectBackupStorage detail param
type ReconnectBackupStorageParamDetail struct {
}

// ReconnectBackupStorageParam ReconnectBackupStorage request param
type ReconnectBackupStorageParam struct {
	BaseParam
	Params ReconnectBackupStorageParamDetail `json:"reconnectBackupStorage"`
}
