// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZStoneLicenseInventoryView ZStoneLicense
type ZStoneLicenseInventoryView struct {
	BaseInfoView
	BaseTimeView
	IssuedTime string `json:"issuedTime,omitempty"`
	ExpiredTime string `json:"expiredTime,omitempty"`
	Expired bool `json:"expired,omitempty"`
	LicenseType string `json:"licenseType,omitempty"`
	ProdInfo string `json:"prodInfo,omitempty"`
	ProductVersion string `json:"productVersion,omitempty"`
	LicenseAttr string `json:"licenseAttr,omitempty"`
	CpuNum int64 `json:"cpuNum,omitempty"`
	UsedCpuNum int64 `json:"usedCpuNum,omitempty"`
	HostNum int64 `json:"hostNum,omitempty"`
	UsedHostNum int64 `json:"usedHostNum,omitempty"`
	Capacity int64 `json:"capacity,omitempty"`
	UsedCapacity int64 `json:"usedCapacity,omitempty"`
}

