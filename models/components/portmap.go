// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Portmap struct {
	V2Entries []PortmapV2Entry `json:"v2_entries,omitempty"`
	V3Entries []PortmapV3Entry `json:"v3_entries,omitempty"`
}

func (o *Portmap) GetV2Entries() []PortmapV2Entry {
	if o == nil {
		return nil
	}
	return o.V2Entries
}

func (o *Portmap) GetV3Entries() []PortmapV3Entry {
	if o == nil {
		return nil
	}
	return o.V3Entries
}
