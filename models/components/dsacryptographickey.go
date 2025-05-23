// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DSACryptographicKey struct {
	G *string `json:"g,omitempty"`
	P *string `json:"p,omitempty"`
	Q *string `json:"q,omitempty"`
	Y *string `json:"y,omitempty"`
}

func (o *DSACryptographicKey) GetG() *string {
	if o == nil {
		return nil
	}
	return o.G
}

func (o *DSACryptographicKey) GetP() *string {
	if o == nil {
		return nil
	}
	return o.P
}

func (o *DSACryptographicKey) GetQ() *string {
	if o == nil {
		return nil
	}
	return o.Q
}

func (o *DSACryptographicKey) GetY() *string {
	if o == nil {
		return nil
	}
	return o.Y
}
