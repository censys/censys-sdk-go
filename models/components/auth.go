// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Auth struct {
	Enabled *bool `json:"enabled,omitempty"`
}

func (o *Auth) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}
