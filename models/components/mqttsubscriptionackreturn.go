// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type MqttSubscriptionAckReturn struct {
	// Raw subscription response value
	Raw *int `json:"raw,omitempty"`
	// Subscription response
	ReturnValue *string `json:"return_value,omitempty"`
}

func (o *MqttSubscriptionAckReturn) GetRaw() *int {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *MqttSubscriptionAckReturn) GetReturnValue() *string {
	if o == nil {
		return nil
	}
	return o.ReturnValue
}
