// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ElasticSearchResultsNodeInfoNodesNodeDataModules struct {
	ClassName      *string  `json:"class_name,omitempty"`
	Desc           *string  `json:"desc,omitempty"`
	ElasticVersion *string  `json:"elastic_version,omitempty"`
	ExtPlugins     []string `json:"ext_plugins,omitempty"`
	HasNativeCtrl  *bool    `json:"has_native_ctrl,omitempty"`
	JavaVersion    *string  `json:"java_version,omitempty"`
	Name           *string  `json:"name,omitempty"`
	Version        *string  `json:"version,omitempty"`
}

func (o *ElasticSearchResultsNodeInfoNodesNodeDataModules) GetClassName() *string {
	if o == nil {
		return nil
	}
	return o.ClassName
}

func (o *ElasticSearchResultsNodeInfoNodesNodeDataModules) GetDesc() *string {
	if o == nil {
		return nil
	}
	return o.Desc
}

func (o *ElasticSearchResultsNodeInfoNodesNodeDataModules) GetElasticVersion() *string {
	if o == nil {
		return nil
	}
	return o.ElasticVersion
}

func (o *ElasticSearchResultsNodeInfoNodesNodeDataModules) GetExtPlugins() []string {
	if o == nil {
		return nil
	}
	return o.ExtPlugins
}

func (o *ElasticSearchResultsNodeInfoNodesNodeDataModules) GetHasNativeCtrl() *bool {
	if o == nil {
		return nil
	}
	return o.HasNativeCtrl
}

func (o *ElasticSearchResultsNodeInfoNodesNodeDataModules) GetJavaVersion() *string {
	if o == nil {
		return nil
	}
	return o.JavaVersion
}

func (o *ElasticSearchResultsNodeInfoNodesNodeDataModules) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ElasticSearchResultsNodeInfoNodesNodeDataModules) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}
