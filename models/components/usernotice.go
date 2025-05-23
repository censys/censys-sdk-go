// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UserNotice struct {
	ExplicitText    *string          `json:"explicit_text,omitempty"`
	NoticeReference *NoticeReference `json:"notice_reference,omitempty"`
}

func (o *UserNotice) GetExplicitText() *string {
	if o == nil {
		return nil
	}
	return o.ExplicitText
}

func (o *UserNotice) GetNoticeReference() *NoticeReference {
	if o == nil {
		return nil
	}
	return o.NoticeReference
}
