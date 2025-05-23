// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type EtcdVersion struct {
	Cluster *string `json:"cluster,omitempty"`
	Server  *string `json:"server,omitempty"`
}

func (o *EtcdVersion) GetCluster() *string {
	if o == nil {
		return nil
	}
	return o.Cluster
}

func (o *EtcdVersion) GetServer() *string {
	if o == nil {
		return nil
	}
	return o.Server
}
