package domain

type EventTriggerDetail struct {
	Id          string      `json:"id"`
	MsgStructId string      `json:"msgStructId"`
	EventDesc   string      `json:"eventDesc"`
	Description string      `json:"description"`
	Sample      interface{} `json:"sample"`
	Segments    []struct {
		Id       *string          `json:"id"`
		Name     string           `json:"name"`
		LongName *string          `json:"longName"`
		Sequence *string          `json:"sequence"`
		Usage    string           `json:"usage"`
		Rpt      string           `json:"rpt"`
		IsGroup  bool             `json:"isGroup"`
		Segments []SegmentForList `json:"segments"`
	} `json:"segments"`
}
