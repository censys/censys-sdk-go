// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CollectionEvent struct {
	AssetChange  *EventAssetChange  `json:"asset_change,omitempty"`
	StatusChange *EventStatusChange `json:"status_change,omitempty"`
}

func (o *CollectionEvent) GetAssetChange() *EventAssetChange {
	if o == nil {
		return nil
	}
	return o.AssetChange
}

func (o *CollectionEvent) GetStatusChange() *EventStatusChange {
	if o == nil {
		return nil
	}
	return o.StatusChange
}
