// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type SapRouterRouterInfo struct {
	ConnectedClientInfo        []SapRouterClientInfo `json:"connected_client_info,omitempty"`
	NumClients                 *int64                `json:"num_clients,omitempty"`
	ParentPid                  *int                  `json:"parent_pid,omitempty"`
	ParentPort                 *int                  `json:"parent_port,omitempty"`
	Pid                        *int                  `json:"pid,omitempty"`
	Port                       *int                  `json:"port,omitempty"`
	RouttabRelativeDirectory   *string               `json:"routtab_relative_directory,omitempty"`
	SapRouterAbsoluteDirectory *string               `json:"sap_router_absolute_directory,omitempty"`
	StartedOn                  *string               `json:"started_on,omitempty"`
}

func (o *SapRouterRouterInfo) GetConnectedClientInfo() []SapRouterClientInfo {
	if o == nil {
		return nil
	}
	return o.ConnectedClientInfo
}

func (o *SapRouterRouterInfo) GetNumClients() *int64 {
	if o == nil {
		return nil
	}
	return o.NumClients
}

func (o *SapRouterRouterInfo) GetParentPid() *int {
	if o == nil {
		return nil
	}
	return o.ParentPid
}

func (o *SapRouterRouterInfo) GetParentPort() *int {
	if o == nil {
		return nil
	}
	return o.ParentPort
}

func (o *SapRouterRouterInfo) GetPid() *int {
	if o == nil {
		return nil
	}
	return o.Pid
}

func (o *SapRouterRouterInfo) GetPort() *int {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SapRouterRouterInfo) GetRouttabRelativeDirectory() *string {
	if o == nil {
		return nil
	}
	return o.RouttabRelativeDirectory
}

func (o *SapRouterRouterInfo) GetSapRouterAbsoluteDirectory() *string {
	if o == nil {
		return nil
	}
	return o.SapRouterAbsoluteDirectory
}

func (o *SapRouterRouterInfo) GetStartedOn() *string {
	if o == nil {
		return nil
	}
	return o.StartedOn
}
