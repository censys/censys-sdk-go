// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type SnmpPhysical struct {
	// 1.3.6.1.2.1.47.1.1.1.1.9 - Firmware revision string
	FirmwareRev *string `json:"firmware_rev,omitempty"`
	// 1.3.6.1.2.1.47.1.1.1.1.8 - Hardware revision string
	HardwareRev *string `json:"hardware_rev,omitempty"`
	// 1.3.6.1.2.1.47.1.1.1.1.12 - Name of mfg
	MfgName *string `json:"mfg_name,omitempty"`
	// 1.3.6.1.2.1.47.1.1.1.1.13 - Model name of component
	ModelName *string `json:"model_name,omitempty"`
	// 1.3.6.1.2.1.47.1.1.1.1.7 - Entity name
	Name *string `json:"name,omitempty"`
	// 1.3.6.1.2.1.47.1.1.1.1.11 - Serial number string
	SerialNum *string `json:"serial_num,omitempty"`
	// 1.3.6.1.2.1.47.1.1.1.1.10 - Software revision string
	SoftwareRev *string `json:"software_rev,omitempty"`
}

func (o *SnmpPhysical) GetFirmwareRev() *string {
	if o == nil {
		return nil
	}
	return o.FirmwareRev
}

func (o *SnmpPhysical) GetHardwareRev() *string {
	if o == nil {
		return nil
	}
	return o.HardwareRev
}

func (o *SnmpPhysical) GetMfgName() *string {
	if o == nil {
		return nil
	}
	return o.MfgName
}

func (o *SnmpPhysical) GetModelName() *string {
	if o == nil {
		return nil
	}
	return o.ModelName
}

func (o *SnmpPhysical) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SnmpPhysical) GetSerialNum() *string {
	if o == nil {
		return nil
	}
	return o.SerialNum
}

func (o *SnmpPhysical) GetSoftwareRev() *string {
	if o == nil {
		return nil
	}
	return o.SoftwareRev
}
