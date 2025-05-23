// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type L2TpAttributeValues struct {
	ErrorCode        *int    `json:"error_code,omitempty"`
	ErrorMeaning     *string `json:"error_meaning,omitempty"`
	ErrorMessage     *string `json:"error_message,omitempty"`
	FirmwareRevision *int    `json:"firmware_revision,omitempty"`
	Hostname         *string `json:"hostname,omitempty"`
	ProtocolRevision *int    `json:"protocol_revision,omitempty"`
	ProtocolVersion  *int    `json:"protocol_version,omitempty"`
	ResultCode       *int    `json:"result_code,omitempty"`
	ResultMeaning    *string `json:"result_meaning,omitempty"`
	VendorName       *string `json:"vendor_name,omitempty"`
	WindowSize       *int    `json:"window_size,omitempty"`
}

func (o *L2TpAttributeValues) GetErrorCode() *int {
	if o == nil {
		return nil
	}
	return o.ErrorCode
}

func (o *L2TpAttributeValues) GetErrorMeaning() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMeaning
}

func (o *L2TpAttributeValues) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *L2TpAttributeValues) GetFirmwareRevision() *int {
	if o == nil {
		return nil
	}
	return o.FirmwareRevision
}

func (o *L2TpAttributeValues) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *L2TpAttributeValues) GetProtocolRevision() *int {
	if o == nil {
		return nil
	}
	return o.ProtocolRevision
}

func (o *L2TpAttributeValues) GetProtocolVersion() *int {
	if o == nil {
		return nil
	}
	return o.ProtocolVersion
}

func (o *L2TpAttributeValues) GetResultCode() *int {
	if o == nil {
		return nil
	}
	return o.ResultCode
}

func (o *L2TpAttributeValues) GetResultMeaning() *string {
	if o == nil {
		return nil
	}
	return o.ResultMeaning
}

func (o *L2TpAttributeValues) GetVendorName() *string {
	if o == nil {
		return nil
	}
	return o.VendorName
}

func (o *L2TpAttributeValues) GetWindowSize() *int {
	if o == nil {
		return nil
	}
	return o.WindowSize
}
