// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CobaltStrikePostEx struct {
	X64 *string `json:"x64,omitempty"`
	X86 *string `json:"x86,omitempty"`
}

func (o *CobaltStrikePostEx) GetX64() *string {
	if o == nil {
		return nil
	}
	return o.X64
}

func (o *CobaltStrikePostEx) GetX86() *string {
	if o == nil {
		return nil
	}
	return o.X86
}
