package abstraction

import "development.zazmic.com/dicom-adapter/adapter-cluster/hl7-v2-def/domain"

type Getter interface {
	GetTriggersEventsByVersion(version string) ([]domain.TriggerEvent, error)

	GetSegmentDetailByVersionAndSegment(version, segmentId string) (domain.Segment, error)

	GetSegmentsByVersionAndTrigger(version, triggerEvent string) (domain.EventTriggerDetail, error)
}
