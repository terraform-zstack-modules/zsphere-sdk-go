package param

type CreateAccessKeyParam struct {
	BaseParam
	Params CreateAccessKeyDetailParams `json:"params"`
}

type CreateAccessKeyDetailParams struct {
	AccountUUID     string  `json:"accountUuid"`
	UserUUID        string  `json:"userUuid"`
	Description     *string `json:"description,omitempty"`
	AccessKeyID     *string `json:"accessKeyID,omitempty"`
	AccessKeySecret *string `json:"accessKeySecret,omitempty"`
}
