// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Dhcpdiscover struct {
	Method *string                     `json:"method,omitempty"`
	Params *DhcpdiscoverResponseParams `json:"params,omitempty"`
}

func (o *Dhcpdiscover) GetMethod() *string {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *Dhcpdiscover) GetParams() *DhcpdiscoverResponseParams {
	if o == nil {
		return nil
	}
	return o.Params
}
