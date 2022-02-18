package domain

type Segment struct {
	Id          string  `json:"id"`
	SegmentId   string  `json:"segmentId"`
	LongName    string  `json:"longName"`
	Description string  `json:"description"`
	Sample      string  `json:"sample"`
	Fields      []Field `json:"fields"`
}
