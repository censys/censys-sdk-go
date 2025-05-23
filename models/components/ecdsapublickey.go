// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type EcdsaPublicKey struct {
	B      *string `json:"b,omitempty"`
	Curve  *string `json:"curve,omitempty"`
	Gx     *string `json:"gx,omitempty"`
	Gy     *string `json:"gy,omitempty"`
	Length *int64  `json:"length,omitempty"`
	N      *string `json:"n,omitempty"`
	P      *string `json:"p,omitempty"`
	Pub    *string `json:"pub,omitempty"`
	X      *string `json:"x,omitempty"`
	Y      *string `json:"y,omitempty"`
}

func (o *EcdsaPublicKey) GetB() *string {
	if o == nil {
		return nil
	}
	return o.B
}

func (o *EcdsaPublicKey) GetCurve() *string {
	if o == nil {
		return nil
	}
	return o.Curve
}

func (o *EcdsaPublicKey) GetGx() *string {
	if o == nil {
		return nil
	}
	return o.Gx
}

func (o *EcdsaPublicKey) GetGy() *string {
	if o == nil {
		return nil
	}
	return o.Gy
}

func (o *EcdsaPublicKey) GetLength() *int64 {
	if o == nil {
		return nil
	}
	return o.Length
}

func (o *EcdsaPublicKey) GetN() *string {
	if o == nil {
		return nil
	}
	return o.N
}

func (o *EcdsaPublicKey) GetP() *string {
	if o == nil {
		return nil
	}
	return o.P
}

func (o *EcdsaPublicKey) GetPub() *string {
	if o == nil {
		return nil
	}
	return o.Pub
}

func (o *EcdsaPublicKey) GetX() *string {
	if o == nil {
		return nil
	}
	return o.X
}

func (o *EcdsaPublicKey) GetY() *string {
	if o == nil {
		return nil
	}
	return o.Y
}
