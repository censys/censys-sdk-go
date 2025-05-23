// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type RecordingCapabilities struct {
	DynamicRecordings *bool   `json:"dynamic_recordings,omitempty"`
	DynamicTracks     *bool   `json:"dynamic_tracks,omitempty"`
	Encoding          *string `json:"encoding,omitempty"`
	MaxRate           *int    `json:"max_rate,omitempty"`
	MaxRecordings     *int    `json:"max_recordings,omitempty"`
	MaxRecordingsJob  *int    `json:"max_recordings_job,omitempty"`
	MaxTotalRate      *int    `json:"max_total_rate,omitempty"`
	Options           *bool   `json:"options,omitempty"`
}

func (o *RecordingCapabilities) GetDynamicRecordings() *bool {
	if o == nil {
		return nil
	}
	return o.DynamicRecordings
}

func (o *RecordingCapabilities) GetDynamicTracks() *bool {
	if o == nil {
		return nil
	}
	return o.DynamicTracks
}

func (o *RecordingCapabilities) GetEncoding() *string {
	if o == nil {
		return nil
	}
	return o.Encoding
}

func (o *RecordingCapabilities) GetMaxRate() *int {
	if o == nil {
		return nil
	}
	return o.MaxRate
}

func (o *RecordingCapabilities) GetMaxRecordings() *int {
	if o == nil {
		return nil
	}
	return o.MaxRecordings
}

func (o *RecordingCapabilities) GetMaxRecordingsJob() *int {
	if o == nil {
		return nil
	}
	return o.MaxRecordingsJob
}

func (o *RecordingCapabilities) GetMaxTotalRate() *int {
	if o == nil {
		return nil
	}
	return o.MaxTotalRate
}

func (o *RecordingCapabilities) GetOptions() *bool {
	if o == nil {
		return nil
	}
	return o.Options
}
