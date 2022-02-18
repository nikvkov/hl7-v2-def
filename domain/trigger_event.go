package domain

type TriggerEvent struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Label       string `json:"label"`
	Description string `json:"description"`
}
