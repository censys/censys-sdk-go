// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Rlogin struct {
	Error           *string `json:"error,omitempty"`
	OperatingSystem *string `json:"operating_system,omitempty"`
	Software        *string `json:"software,omitempty"`
	SoftwareVersion *string `json:"software_version,omitempty"`
}

func (o *Rlogin) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *Rlogin) GetOperatingSystem() *string {
	if o == nil {
		return nil
	}
	return o.OperatingSystem
}

func (o *Rlogin) GetSoftware() *string {
	if o == nil {
		return nil
	}
	return o.Software
}

func (o *Rlogin) GetSoftwareVersion() *string {
	if o == nil {
		return nil
	}
	return o.SoftwareVersion
}
