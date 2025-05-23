// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Socks struct {
	NoAuthenticationRequired     *bool   `json:"no_authentication_required,omitempty"`
	PreferredAuthentication      *string `json:"preferred_authentication,omitempty"`
	PreferredAuthenticationValue *int    `json:"preferred_authentication_value,omitempty"`
	SocksVersion                 *int64  `json:"socks_version,omitempty"`
}

func (o *Socks) GetNoAuthenticationRequired() *bool {
	if o == nil {
		return nil
	}
	return o.NoAuthenticationRequired
}

func (o *Socks) GetPreferredAuthentication() *string {
	if o == nil {
		return nil
	}
	return o.PreferredAuthentication
}

func (o *Socks) GetPreferredAuthenticationValue() *int {
	if o == nil {
		return nil
	}
	return o.PreferredAuthenticationValue
}

func (o *Socks) GetSocksVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.SocksVersion
}
