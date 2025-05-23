// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Worker struct {
	ClientID  *string `json:"client_id,omitempty"`
	Fd        *string `json:"fd,omitempty"`
	Functions *string `json:"functions,omitempty"`
	IP        *string `json:"ip,omitempty"`
}

func (o *Worker) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *Worker) GetFd() *string {
	if o == nil {
		return nil
	}
	return o.Fd
}

func (o *Worker) GetFunctions() *string {
	if o == nil {
		return nil
	}
	return o.Functions
}

func (o *Worker) GetIP() *string {
	if o == nil {
		return nil
	}
	return o.IP
}
