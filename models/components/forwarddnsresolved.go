// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ForwardDNSResolved struct {
	Name        *string `json:"name,omitempty"`
	ResolveTime *string `json:"resolve_time,omitempty"`
}

func (o *ForwardDNSResolved) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ForwardDNSResolved) GetResolveTime() *string {
	if o == nil {
		return nil
	}
	return o.ResolveTime
}
