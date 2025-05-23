// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type IpmiCommandPayloadPackedNetFnNetFn struct {
	// True if the least-significant bit is zero
	IsRequest *bool `json:"is_request,omitempty"`
	// True if the least-significant bit is one
	IsResponse *bool `json:"is_response,omitempty"`
	// The human-readable name of the NetFn
	Name *string `json:"name,omitempty"`
	// The raw value of the NetFn (6 bits, least significant indicates request/response)
	Raw *int `json:"raw,omitempty"`
	// The normalized value of the NetFn (i.e. raw & 0xfe, so it is always even)
	Value *int `json:"value,omitempty"`
}

func (o *IpmiCommandPayloadPackedNetFnNetFn) GetIsRequest() *bool {
	if o == nil {
		return nil
	}
	return o.IsRequest
}

func (o *IpmiCommandPayloadPackedNetFnNetFn) GetIsResponse() *bool {
	if o == nil {
		return nil
	}
	return o.IsResponse
}

func (o *IpmiCommandPayloadPackedNetFnNetFn) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *IpmiCommandPayloadPackedNetFnNetFn) GetRaw() *int {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *IpmiCommandPayloadPackedNetFnNetFn) GetValue() *int {
	if o == nil {
		return nil
	}
	return o.Value
}
