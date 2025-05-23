// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type X11 struct {
	RefusalReason          *string `json:"refusal_reason,omitempty"`
	RequiresAuthentication *bool   `json:"requires_authentication,omitempty"`
	Vendor                 *string `json:"vendor,omitempty"`
	Version                *string `json:"version,omitempty"`
}

func (o *X11) GetRefusalReason() *string {
	if o == nil {
		return nil
	}
	return o.RefusalReason
}

func (o *X11) GetRequiresAuthentication() *bool {
	if o == nil {
		return nil
	}
	return o.RequiresAuthentication
}

func (o *X11) GetVendor() *string {
	if o == nil {
		return nil
	}
	return o.Vendor
}

func (o *X11) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}
