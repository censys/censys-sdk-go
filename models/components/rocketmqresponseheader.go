// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type RocketmqResponseHeader struct {
	Code                    *int64  `json:"code,omitempty"`
	Flag                    *int64  `json:"flag,omitempty"`
	Language                *string `json:"language,omitempty"`
	Opaque                  *int64  `json:"opaque,omitempty"`
	SerializeTypeCurrentRPC *string `json:"serialize_type_current_rpc,omitempty"`
}

func (o *RocketmqResponseHeader) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *RocketmqResponseHeader) GetFlag() *int64 {
	if o == nil {
		return nil
	}
	return o.Flag
}

func (o *RocketmqResponseHeader) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *RocketmqResponseHeader) GetOpaque() *int64 {
	if o == nil {
		return nil
	}
	return o.Opaque
}

func (o *RocketmqResponseHeader) GetSerializeTypeCurrentRPC() *string {
	if o == nil {
		return nil
	}
	return o.SerializeTypeCurrentRPC
}
