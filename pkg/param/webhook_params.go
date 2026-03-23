// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateWebhookParamDetail CreateWebhook detail param
type CreateWebhookParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Url string `json:"url" validate:"required"`
	Type string `json:"type" validate:"required"`
	Opaque *string `json:"opaque,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateWebhookParam CreateWebhook request param
type CreateWebhookParam struct {
	BaseParam
	Params CreateWebhookParamDetail `json:"params"`
}
// UpdateWebhookParamDetail UpdateWebhook detail param
type UpdateWebhookParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Url *string `json:"url,omitempty"`
	Type *string `json:"type,omitempty"`
	Opaque *string `json:"opaque,omitempty"`
}

// UpdateWebhookParam UpdateWebhook request param
type UpdateWebhookParam struct {
	BaseParam
	Params UpdateWebhookParamDetail `json:"updateWebhook"`
}
// DeleteWebhookParamDetail DeleteWebhook detail param
type DeleteWebhookParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteWebhookParam DeleteWebhook request param
type DeleteWebhookParam struct {
	BaseParam
	Params DeleteWebhookParamDetail `json:"deleteWebhook"`
}
