// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Onvif struct {
	Hostname *OnvifHostname `json:"hostname,omitempty"`
	Services []OnvifService `json:"services,omitempty"`
}

func (o *Onvif) GetHostname() *OnvifHostname {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *Onvif) GetServices() []OnvifService {
	if o == nil {
		return nil
	}
	return o.Services
}
