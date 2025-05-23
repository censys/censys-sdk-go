// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UpnpDevice struct {
	DeviceType   *string `json:"device_type,omitempty"`
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Censys-generated IDs representing a device tree
	ID               *int          `json:"id,omitempty"`
	Manufacturer     *string       `json:"manufacturer,omitempty"`
	ManufacturerURL  *string       `json:"manufacturer_url,omitempty"`
	ModelDescription *string       `json:"model_description,omitempty"`
	ModelName        *string       `json:"model_name,omitempty"`
	ModelNumber      *string       `json:"model_number,omitempty"`
	ModelURL         *string       `json:"model_url,omitempty"`
	ParentID         *int          `json:"parent_id,omitempty"`
	PresentationURL  *string       `json:"presentation_url,omitempty"`
	SerialNumber     *string       `json:"serial_number,omitempty"`
	ServiceList      []UpnpService `json:"service_list,omitempty"`
	Udn              *string       `json:"udn,omitempty"`
	Upc              *string       `json:"upc,omitempty"`
}

func (o *UpnpDevice) GetDeviceType() *string {
	if o == nil {
		return nil
	}
	return o.DeviceType
}

func (o *UpnpDevice) GetFriendlyName() *string {
	if o == nil {
		return nil
	}
	return o.FriendlyName
}

func (o *UpnpDevice) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpnpDevice) GetManufacturer() *string {
	if o == nil {
		return nil
	}
	return o.Manufacturer
}

func (o *UpnpDevice) GetManufacturerURL() *string {
	if o == nil {
		return nil
	}
	return o.ManufacturerURL
}

func (o *UpnpDevice) GetModelDescription() *string {
	if o == nil {
		return nil
	}
	return o.ModelDescription
}

func (o *UpnpDevice) GetModelName() *string {
	if o == nil {
		return nil
	}
	return o.ModelName
}

func (o *UpnpDevice) GetModelNumber() *string {
	if o == nil {
		return nil
	}
	return o.ModelNumber
}

func (o *UpnpDevice) GetModelURL() *string {
	if o == nil {
		return nil
	}
	return o.ModelURL
}

func (o *UpnpDevice) GetParentID() *int {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *UpnpDevice) GetPresentationURL() *string {
	if o == nil {
		return nil
	}
	return o.PresentationURL
}

func (o *UpnpDevice) GetSerialNumber() *string {
	if o == nil {
		return nil
	}
	return o.SerialNumber
}

func (o *UpnpDevice) GetServiceList() []UpnpService {
	if o == nil {
		return nil
	}
	return o.ServiceList
}

func (o *UpnpDevice) GetUdn() *string {
	if o == nil {
		return nil
	}
	return o.Udn
}

func (o *UpnpDevice) GetUpc() *string {
	if o == nil {
		return nil
	}
	return o.Upc
}
