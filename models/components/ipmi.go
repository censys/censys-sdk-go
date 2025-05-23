// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Ipmi struct {
	Capabilities   *IpmiCapabilities   `json:"capabilities,omitempty"`
	CommandPayload *IpmiCommandPayload `json:"command_payload,omitempty"`
	// The raw data returned by the server
	Raw           *string            `json:"raw,omitempty"`
	RmcpHeader    *IpmiRMCPHeader    `json:"rmcp_header,omitempty"`
	SessionHeader *IpmiSessionHeader `json:"session_header,omitempty"`
}

func (o *Ipmi) GetCapabilities() *IpmiCapabilities {
	if o == nil {
		return nil
	}
	return o.Capabilities
}

func (o *Ipmi) GetCommandPayload() *IpmiCommandPayload {
	if o == nil {
		return nil
	}
	return o.CommandPayload
}

func (o *Ipmi) GetRaw() *string {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *Ipmi) GetRmcpHeader() *IpmiRMCPHeader {
	if o == nil {
		return nil
	}
	return o.RmcpHeader
}

func (o *Ipmi) GetSessionHeader() *IpmiSessionHeader {
	if o == nil {
		return nil
	}
	return o.SessionHeader
}
