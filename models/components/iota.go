// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Iota struct {
	V0Info *NodeInfoV0 `json:"v0_info,omitempty"`
	V1Info *NodeInfoV1 `json:"v1_info,omitempty"`
	V2Info *NodeInfoV2 `json:"v2_info,omitempty"`
}

func (o *Iota) GetV0Info() *NodeInfoV0 {
	if o == nil {
		return nil
	}
	return o.V0Info
}

func (o *Iota) GetV1Info() *NodeInfoV1 {
	if o == nil {
		return nil
	}
	return o.V1Info
}

func (o *Iota) GetV2Info() *NodeInfoV2 {
	if o == nil {
		return nil
	}
	return o.V2Info
}
