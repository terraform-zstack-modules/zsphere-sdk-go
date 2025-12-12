package view

type AccessKeyInventoryView struct {
	BaseTimeView
	UUID            string `json:"uuid"`
	Description     string `json:"description"`
	AccountUUID     string `json:"accountUuid"`
	UserUUID        string `json:"userUuid"`
	AccessKeyID     string `json:"accessKeyID"`
	AccessKeySecret string `json:"accessKeySecret"`
	State           string `json:"state"`
}
