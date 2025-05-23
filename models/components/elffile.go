// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ElfFile struct {
	Class   *string `json:"class,omitempty"`
	Data    *string `json:"data,omitempty"`
	Machine *string `json:"machine,omitempty"`
	OsAbi   *string `json:"os_abi,omitempty"`
	Type    *string `json:"type,omitempty"`
}

func (o *ElfFile) GetClass() *string {
	if o == nil {
		return nil
	}
	return o.Class
}

func (o *ElfFile) GetData() *string {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *ElfFile) GetMachine() *string {
	if o == nil {
		return nil
	}
	return o.Machine
}

func (o *ElfFile) GetOsAbi() *string {
	if o == nil {
		return nil
	}
	return o.OsAbi
}

func (o *ElfFile) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}
