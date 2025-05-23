// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type VentriloHeader struct {
	Cmd         *int `json:"cmd,omitempty"`
	DataKey     *int `json:"data_key,omitempty"`
	HeaderKey   *int `json:"header_key,omitempty"`
	ID          *int `json:"id,omitempty"`
	TotalLength *int `json:"total_length,omitempty"`
}

func (o *VentriloHeader) GetCmd() *int {
	if o == nil {
		return nil
	}
	return o.Cmd
}

func (o *VentriloHeader) GetDataKey() *int {
	if o == nil {
		return nil
	}
	return o.DataKey
}

func (o *VentriloHeader) GetHeaderKey() *int {
	if o == nil {
		return nil
	}
	return o.HeaderKey
}

func (o *VentriloHeader) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *VentriloHeader) GetTotalLength() *int {
	if o == nil {
		return nil
	}
	return o.TotalLength
}
