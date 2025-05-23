// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type MssqlPreloginOptions struct {
	EncryptMode     *string                            `json:"encrypt_mode,omitempty"`
	FedAuthRequired *bool                              `json:"fed_auth_required,omitempty"`
	Instance        *string                            `json:"instance,omitempty"`
	Mars            *bool                              `json:"mars,omitempty"`
	Nonce           *string                            `json:"nonce,omitempty"`
	ServerVersion   *MssqlPreloginOptionsServerVersion `json:"server_version,omitempty"`
	ThreadID        *int                               `json:"thread_id,omitempty"`
	TraceID         *string                            `json:"trace_id,omitempty"`
	Unknown         map[string]string                  `json:"unknown,omitempty"`
}

func (o *MssqlPreloginOptions) GetEncryptMode() *string {
	if o == nil {
		return nil
	}
	return o.EncryptMode
}

func (o *MssqlPreloginOptions) GetFedAuthRequired() *bool {
	if o == nil {
		return nil
	}
	return o.FedAuthRequired
}

func (o *MssqlPreloginOptions) GetInstance() *string {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *MssqlPreloginOptions) GetMars() *bool {
	if o == nil {
		return nil
	}
	return o.Mars
}

func (o *MssqlPreloginOptions) GetNonce() *string {
	if o == nil {
		return nil
	}
	return o.Nonce
}

func (o *MssqlPreloginOptions) GetServerVersion() *MssqlPreloginOptionsServerVersion {
	if o == nil {
		return nil
	}
	return o.ServerVersion
}

func (o *MssqlPreloginOptions) GetThreadID() *int {
	if o == nil {
		return nil
	}
	return o.ThreadID
}

func (o *MssqlPreloginOptions) GetTraceID() *string {
	if o == nil {
		return nil
	}
	return o.TraceID
}

func (o *MssqlPreloginOptions) GetUnknown() map[string]string {
	if o == nil {
		return nil
	}
	return o.Unknown
}
