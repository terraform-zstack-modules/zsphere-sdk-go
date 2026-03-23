// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LicenseHistoryInventoryView LicenseHistory
type LicenseHistoryInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	Quota int64 `json:"quota,omitempty"`
	QuotaType string `json:"quotaType,omitempty"`
	ExpiredDate string `json:"expiredDate,omitempty"`
	IssuedDate string `json:"issuedDate,omitempty"`
	UploadDate string `json:"uploadDate,omitempty"`
	LicenseType string `json:"licenseType,omitempty"`
	ProdInfo string `json:"prodInfo,omitempty"`
	UserName string `json:"userName,omitempty"`
	Hash string `json:"hash,omitempty"`
	Source string `json:"source,omitempty"`
	ManagementNodeUuid string `json:"managementNodeUuid,omitempty"`
	MergedTo int64 `json:"mergedTo,omitempty"`
	Expired bool `json:"expired,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	HostNum int `json:"hostNum,omitempty"`
	VmNum int `json:"vmNum,omitempty"`
	Capacity int `json:"capacity,omitempty"`
}

// GetLicenseRecordsView GetLicenseRecords
type GetLicenseRecordsView struct {
	Inventories []LicenseHistoryInventoryView `json:"inventories,omitempty"`
	Total int64 `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

