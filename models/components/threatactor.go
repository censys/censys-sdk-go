// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ThreatActor struct {
	AllNames        []string `json:"all_names,omitempty"`
	ID              *string  `json:"id,omitempty"`
	MalpediaGroupID *string  `json:"malpedia_group_id,omitempty"`
	MitreGroupID    *string  `json:"mitre_group_id,omitempty"`
	PrimaryName     *string  `json:"primary_name,omitempty"`
}

func (o *ThreatActor) GetAllNames() []string {
	if o == nil {
		return nil
	}
	return o.AllNames
}

func (o *ThreatActor) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ThreatActor) GetMalpediaGroupID() *string {
	if o == nil {
		return nil
	}
	return o.MalpediaGroupID
}

func (o *ThreatActor) GetMitreGroupID() *string {
	if o == nil {
		return nil
	}
	return o.MitreGroupID
}

func (o *ThreatActor) GetPrimaryName() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryName
}
