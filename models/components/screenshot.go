// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Screenshot struct {
	ExtractedText *string `json:"extracted_text,omitempty"`
	Handle        *string `json:"handle,omitempty"`
}

func (o *Screenshot) GetExtractedText() *string {
	if o == nil {
		return nil
	}
	return o.ExtractedText
}

func (o *Screenshot) GetHandle() *string {
	if o == nil {
		return nil
	}
	return o.Handle
}
