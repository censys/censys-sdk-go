// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CrudCreateInputBody struct {
	// description of the collection
	Description *string `json:"description,omitempty"`
	// name of the collection
	Name string `json:"name"`
	// query string to search upon to build the collection
	Query string `json:"query"`
}

func (o *CrudCreateInputBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CrudCreateInputBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CrudCreateInputBody) GetQuery() string {
	if o == nil {
		return ""
	}
	return o.Query
}
