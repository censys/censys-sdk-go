// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ResponseEnvelopeCertificateAsset struct {
	Result *CertificateAsset `json:"result,omitempty"`
}

func (o *ResponseEnvelopeCertificateAsset) GetResult() *CertificateAsset {
	if o == nil {
		return nil
	}
	return o.Result
}
