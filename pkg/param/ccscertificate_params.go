// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// AddCCSCertificateParamDetail AddCCSCertificate detail param
type AddCCSCertificateParamDetail struct {
	Certificate string `json:"certificate" validate:"required"`
}

// AddCCSCertificateParam AddCCSCertificate request param
type AddCCSCertificateParam struct {
	BaseParam
	Params AddCCSCertificateParamDetail `json:"addCCSCertificate"`
}
// DeleteCCSCertificateParamDetail DeleteCCSCertificate detail param
type DeleteCCSCertificateParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteCCSCertificateParam DeleteCCSCertificate request param
type DeleteCCSCertificateParam struct {
	BaseParam
	Params DeleteCCSCertificateParamDetail `json:"deleteCCSCertificate"`
}
