// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Graphql struct {
	Response              *string `json:"response,omitempty"`
	SupportsIntrospection *bool   `json:"supports_introspection,omitempty"`
}

func (o *Graphql) GetResponse() *string {
	if o == nil {
		return nil
	}
	return o.Response
}

func (o *Graphql) GetSupportsIntrospection() *bool {
	if o == nil {
		return nil
	}
	return o.SupportsIntrospection
}
