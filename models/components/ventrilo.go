// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Ventrilo struct {
	Attributes []string          `json:"attributes,omitempty"`
	Messages   []VentriloMessage `json:"messages,omitempty"`
}

func (o *Ventrilo) GetAttributes() []string {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *Ventrilo) GetMessages() []VentriloMessage {
	if o == nil {
		return nil
	}
	return o.Messages
}
