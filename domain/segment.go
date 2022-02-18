package domain

type SegmentForList struct {
	Id       string      `json:"id"`
	Name     string      `json:"name"`
	LongName string      `json:"longName"`
	Sequence string      `json:"sequence"`
	Usage    string      `json:"usage"`
	Rpt      string      `json:"rpt"`
	IsGroup  bool        `json:"isGroup"`
	Segments interface{} `json:"segments"`
}
