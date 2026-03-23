// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteCdpPolicyParamDetail DeleteCdpPolicy detail param
type DeleteCdpPolicyParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteCdpPolicyParam DeleteCdpPolicy request param
type DeleteCdpPolicyParam struct {
	BaseParam
	Params DeleteCdpPolicyParamDetail `json:"deleteCdpPolicy"`
}
// CreateCdpPolicyParamDetail CreateCdpPolicy detail param
type CreateCdpPolicyParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	RetentionTimePerDay *int `json:"retentionTimePerDay,omitempty"`
	HourlyRpSinceDay *int `json:"hourlyRpSinceDay,omitempty"`
	DailyRpSinceDay *int `json:"dailyRpSinceDay,omitempty"`
	ExpireTimeInDay *int `json:"expireTimeInDay,omitempty"`
	FullBackupIntervalInDay *int `json:"fullBackupIntervalInDay,omitempty"`
	RecoveryPointPerSecond int `json:"recoveryPointPerSecond" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateCdpPolicyParam CreateCdpPolicy request param
type CreateCdpPolicyParam struct {
	BaseParam
	Params CreateCdpPolicyParamDetail `json:"params"`
}
// UpdateCdpPolicyParamDetail UpdateCdpPolicy detail param
type UpdateCdpPolicyParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	RetentionTimePerDay *int `json:"retentionTimePerDay,omitempty"`
	HourlyRpSinceDay *int `json:"hourlyRpSinceDay,omitempty"`
	DailyRpSinceDay *int `json:"dailyRpSinceDay,omitempty"`
	ExpireTimeInDay *int `json:"expireTimeInDay,omitempty"`
	FullBackupIntervalInDay *int `json:"fullBackupIntervalInDay,omitempty"`
	RecoveryPointPerSecond *int `json:"recoveryPointPerSecond,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// UpdateCdpPolicyParam UpdateCdpPolicy request param
type UpdateCdpPolicyParam struct {
	BaseParam
	Params UpdateCdpPolicyParamDetail `json:"updateCdpPolicy"`
}
