// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type GeneralSubtreeIP struct {
	// The first IP address in the range.
	Begin *string `json:"begin,omitempty"`
	// The CIDR specifying the subtree.
	Cidr *string `json:"cidr,omitempty"`
	// The last IP address in the range.
	End *string `json:"end,omitempty"`
	// The subnet mask of the CIDR.
	Mask *string `json:"mask,omitempty"`
}

func (o *GeneralSubtreeIP) GetBegin() *string {
	if o == nil {
		return nil
	}
	return o.Begin
}

func (o *GeneralSubtreeIP) GetCidr() *string {
	if o == nil {
		return nil
	}
	return o.Cidr
}

func (o *GeneralSubtreeIP) GetEnd() *string {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *GeneralSubtreeIP) GetMask() *string {
	if o == nil {
		return nil
	}
	return o.Mask
}
