// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type MediaProfile struct {
	MaxProfileCount *int `json:"max_profile_count,omitempty"`
}

func (o *MediaProfile) GetMaxProfileCount() *int {
	if o == nil {
		return nil
	}
	return o.MaxProfileCount
}
