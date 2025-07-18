// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/censys/censys-sdk-go/internal/utils"
	"github.com/censys/censys-sdk-go/models/components"
	"time"
)

type V3GlobaldataAssetWebpropertyGlobals struct {
	OrganizationID *string `queryParam:"style=form,explode=true,name=organization_id"`
}

func (o *V3GlobaldataAssetWebpropertyGlobals) GetOrganizationID() *string {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

type V3GlobaldataAssetWebpropertyRequest struct {
	// The ID of a Censys organization to associate the request with. See the [Getting Started docs](https://docs.censys.com/reference/get-started#step-3-set-your-organization-id) for more information.
	OrganizationID *string `queryParam:"style=form,explode=false,name=organization_id"`
	// A web property identifier.
	WebpropertyID string `pathParam:"style=simple,explode=false,name=webproperty_id"`
	// RFC3339 Timestamp to view a webproperty at a specific point in time. Must be a valid RFC3339 string. Ensure that you suffix the date with T00:00:00Z or a specific time
	AtTime *time.Time `queryParam:"style=form,explode=false,name=at_time"`
}

func (v V3GlobaldataAssetWebpropertyRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V3GlobaldataAssetWebpropertyRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V3GlobaldataAssetWebpropertyRequest) GetOrganizationID() *string {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

func (o *V3GlobaldataAssetWebpropertyRequest) GetWebpropertyID() string {
	if o == nil {
		return ""
	}
	return o.WebpropertyID
}

func (o *V3GlobaldataAssetWebpropertyRequest) GetAtTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.AtTime
}

type V3GlobaldataAssetWebpropertyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ResponseEnvelopeWebpropertyAsset *components.ResponseEnvelopeWebpropertyAsset
	Headers                          map[string][]string
}

func (o *V3GlobaldataAssetWebpropertyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *V3GlobaldataAssetWebpropertyResponse) GetResponseEnvelopeWebpropertyAsset() *components.ResponseEnvelopeWebpropertyAsset {
	if o == nil {
		return nil
	}
	return o.ResponseEnvelopeWebpropertyAsset
}

func (o *V3GlobaldataAssetWebpropertyResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
