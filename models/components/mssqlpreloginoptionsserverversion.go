// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type MssqlPreloginOptionsServerVersion struct {
	BuildNumber *int `json:"build_number,omitempty"`
	Major       *int `json:"major,omitempty"`
	Minor       *int `json:"minor,omitempty"`
}

func (o *MssqlPreloginOptionsServerVersion) GetBuildNumber() *int {
	if o == nil {
		return nil
	}
	return o.BuildNumber
}

func (o *MssqlPreloginOptionsServerVersion) GetMajor() *int {
	if o == nil {
		return nil
	}
	return o.Major
}

func (o *MssqlPreloginOptionsServerVersion) GetMinor() *int {
	if o == nil {
		return nil
	}
	return o.Minor
}
