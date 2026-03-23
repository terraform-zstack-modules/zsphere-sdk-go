// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteCertificateParamDetail DeleteCertificate detail param
type DeleteCertificateParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteCertificateParam DeleteCertificate request param
type DeleteCertificateParam struct {
	BaseParam
	Params DeleteCertificateParamDetail `json:"deleteCertificate"`
}
// UpdateCertificateParamDetail UpdateCertificate detail param
type UpdateCertificateParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// UpdateCertificateParam UpdateCertificate request param
type UpdateCertificateParam struct {
	BaseParam
	Params UpdateCertificateParamDetail `json:"updateCertificate"`
}
// CreateCertificateParamDetail CreateCertificate detail param
type CreateCertificateParamDetail struct {
	Name string `json:"name" validate:"required"`
	Certificate string `json:"certificate" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateCertificateParam CreateCertificate request param
type CreateCertificateParam struct {
	BaseParam
	Params CreateCertificateParamDetail `json:"params"`
}
