package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"development.zazmic.com/dicom-adapter/adapter-cluster/hl7-v2-def/domain"
)

const (
	baseUrl                     = "https://hl7-definition.caristix.com/v2-api/1"
	triggersUrl                 = baseUrl + "/%s/TriggerEvents"
	segmentUrl                  = baseUrl + "/%s/Segments/%s"
	segmentsByVersionAndTrigger = baseUrl + "/%s/TriggerEvents/%s"
)

type httpclient struct {
}

func NewHttpClient() *httpclient {
	return &httpclient{}
}

func (h *httpclient) GetTriggersEventsByVersion(version string) ([]domain.TriggerEvent, error) {
	if version == "" {
		version = domain.HL7v2version24
	}
	url := fmt.Sprintf(triggersUrl, version)
	body, err := h.getBody(url)
	if err != nil {
		return nil, err
	}
	var result []domain.TriggerEvent
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (h *httpclient) GetSegmentDetailByVersionAndSegment(version, segmentId string) (domain.Segment, error) {
	if version == "" {
		version = domain.HL7v2version24
	}
	if segmentId == "" {
		segmentId = "MSH"
	}

	url := fmt.Sprintf(segmentUrl, version, segmentId)
	body, err := h.getBody(url)
	if err != nil {
		return domain.Segment{}, err
	}
	var result domain.Segment
	if err := json.Unmarshal(body, &result); err != nil {
		return domain.Segment{}, err
	}
	return result, nil
}

func (h *httpclient) GetSegmentsByVersionAndTrigger(version, triggerEvent string) (domain.EventTriggerDetail, error) {
	if version == "" {
		version = domain.HL7v2version24
	}
	if triggerEvent == "" {
		triggerEvent = "ADT_A01"
	}

	url := fmt.Sprintf(segmentsByVersionAndTrigger, version, triggerEvent)
	body, err := h.getBody(url)
	if err != nil {
		return domain.EventTriggerDetail{}, err
	}
	var result domain.EventTriggerDetail
	if err := json.Unmarshal(body, &result); err != nil {
		return domain.EventTriggerDetail{}, err
	}
	return result, nil
}

func (h *httpclient) getBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
