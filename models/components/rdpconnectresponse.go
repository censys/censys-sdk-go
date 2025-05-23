// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type RdpConnectResponse struct {
	ConnectID        *int                 `json:"connect_id,omitempty"`
	DomainParameters *RdpDomainParameters `json:"domain_parameters,omitempty"`
}

func (o *RdpConnectResponse) GetConnectID() *int {
	if o == nil {
		return nil
	}
	return o.ConnectID
}

func (o *RdpConnectResponse) GetDomainParameters() *RdpDomainParameters {
	if o == nil {
		return nil
	}
	return o.DomainParameters
}
