// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type PptpVersion struct {
	Major *int `json:"major,omitempty"`
	Minor *int `json:"minor,omitempty"`
}

func (o *PptpVersion) GetMajor() *int {
	if o == nil {
		return nil
	}
	return o.Major
}

func (o *PptpVersion) GetMinor() *int {
	if o == nil {
		return nil
	}
	return o.Minor
}
