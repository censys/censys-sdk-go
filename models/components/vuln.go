// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type VulnRiskSource string

const (
	VulnRiskSourceUnknown VulnRiskSource = ""
	VulnRiskSourceCensys  VulnRiskSource = "censys"
	VulnRiskSourceCve     VulnRiskSource = "cve"
)

func (e VulnRiskSource) ToPointer() *VulnRiskSource {
	return &e
}
func (e *VulnRiskSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "censys":
		fallthrough
	case "cve":
		*e = VulnRiskSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VulnRiskSource: %v", v)
	}
}

type VulnSeverity string

const (
	VulnSeverityUnknown  VulnSeverity = ""
	VulnSeverityLow      VulnSeverity = "low"
	VulnSeverityMedium   VulnSeverity = "medium"
	VulnSeverityHigh     VulnSeverity = "high"
	VulnSeverityCritical VulnSeverity = "critical"
)

func (e VulnSeverity) ToPointer() *VulnSeverity {
	return &e
}
func (e *VulnSeverity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "low":
		fallthrough
	case "medium":
		fallthrough
	case "high":
		fallthrough
	case "critical":
		*e = VulnSeverity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VulnSeverity: %v", v)
	}
}

type VulnSource string

const (
	VulnSourceUnknown    VulnSource = ""
	VulnSourceCensys     VulnSource = "censys"
	VulnSourceRecog      VulnSource = "recog"
	VulnSourceWappalyzer VulnSource = "wappalyzer"
	VulnSourceThirdParty VulnSource = "third_party"
)

func (e VulnSource) ToPointer() *VulnSource {
	return &e
}
func (e *VulnSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "censys":
		fallthrough
	case "recog":
		fallthrough
	case "wappalyzer":
		fallthrough
	case "third_party":
		*e = VulnSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VulnSource: %v", v)
	}
}

type Vuln struct {
	Confidence *float64        `json:"confidence,omitempty"`
	Cwes       []Cwe           `json:"cwes,omitempty"`
	Evidence   []Evidence      `json:"evidence,omitempty"`
	ID         *string         `json:"id,omitempty"`
	Kev        []Kev           `json:"kev,omitempty"`
	Metrics    *Metrics        `json:"metrics,omitempty"`
	Name       *string         `json:"name,omitempty"`
	RiskSource *VulnRiskSource `json:"risk_source,omitempty"`
	Severity   *VulnSeverity   `json:"severity,omitempty"`
	Source     *VulnSource     `json:"source,omitempty"`
	Year       *int            `json:"year,omitempty"`
}

func (o *Vuln) GetConfidence() *float64 {
	if o == nil {
		return nil
	}
	return o.Confidence
}

func (o *Vuln) GetCwes() []Cwe {
	if o == nil {
		return nil
	}
	return o.Cwes
}

func (o *Vuln) GetEvidence() []Evidence {
	if o == nil {
		return nil
	}
	return o.Evidence
}

func (o *Vuln) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Vuln) GetKev() []Kev {
	if o == nil {
		return nil
	}
	return o.Kev
}

func (o *Vuln) GetMetrics() *Metrics {
	if o == nil {
		return nil
	}
	return o.Metrics
}

func (o *Vuln) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Vuln) GetRiskSource() *VulnRiskSource {
	if o == nil {
		return nil
	}
	return o.RiskSource
}

func (o *Vuln) GetSeverity() *VulnSeverity {
	if o == nil {
		return nil
	}
	return o.Severity
}

func (o *Vuln) GetSource() *VulnSource {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *Vuln) GetYear() *int {
	if o == nil {
		return nil
	}
	return o.Year
}
