package hl7_v2_def

import (
	"development.zazmic.com/dicom-adapter/adapter-cluster/hl7-v2-def/domain"
	"development.zazmic.com/dicom-adapter/adapter-cluster/hl7-v2-def/internal/abstraction"
	"development.zazmic.com/dicom-adapter/adapter-cluster/hl7-v2-def/internal/client"
)

type getter struct {
	g abstraction.Getter
}

func NewGetter() *getter {
	return &getter{
		g: client.NewHttpClient(),
	}
}

func (gt *getter) GetTriggerEventList(version string) ([]domain.TriggerEvent, error) {
	return gt.g.GetTriggersEventsByVersion(version)
}

func (gt *getter) GetSegmentListByVersionAndTriggerEvent(version, triggerEvent string) (domain.EventTriggerDetail, error) {
	return gt.g.GetSegmentsByVersionAndTrigger(version, triggerEvent)
}

func (gt *getter) GetSegmentDetailByVersion(version, segment string) (domain.Segment, error) {
	return gt.g.GetSegmentDetailByVersionAndSegment(version, segment)
}
