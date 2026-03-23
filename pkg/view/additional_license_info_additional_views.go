// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AdditionalLicenseInfoView AdditionalLicenseInfo
type AdditionalLicenseInfoView struct {
	Type string `json:"type,omitempty"`
	PrimaryLicenseInfo string `json:"primaryLicenseInfo,omitempty"`
	AppId string `json:"appId,omitempty"`
	Path string `json:"path,omitempty"`
	Info string `json:"info,omitempty"`
	KeyId string `json:"keyId,omitempty"`
}

