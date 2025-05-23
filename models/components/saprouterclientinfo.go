// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type SapRouterClientInfo struct {
	Connected   *bool   `json:"connected,omitempty"`
	ConnectedOn *string `json:"connected_on,omitempty"`
	ID          *int    `json:"id,omitempty"`
	Routed      *bool   `json:"routed,omitempty"`
	Service     *string `json:"service,omitempty"`
	Traced      *bool   `json:"traced,omitempty"`
}

func (o *SapRouterClientInfo) GetConnected() *bool {
	if o == nil {
		return nil
	}
	return o.Connected
}

func (o *SapRouterClientInfo) GetConnectedOn() *string {
	if o == nil {
		return nil
	}
	return o.ConnectedOn
}

func (o *SapRouterClientInfo) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SapRouterClientInfo) GetRouted() *bool {
	if o == nil {
		return nil
	}
	return o.Routed
}

func (o *SapRouterClientInfo) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *SapRouterClientInfo) GetTraced() *bool {
	if o == nil {
		return nil
	}
	return o.Traced
}
