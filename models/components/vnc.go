// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Vnc struct {
	// If server terminates handshake, the reason offered (if any)
	ConnectionFailedReason *string `json:"connection_failed_reason,omitempty"`
	// Desktop name provided by the server, capped at 255 bytes
	DesktopName   *string      `json:"desktop_name,omitempty"`
	PixelEncoding *VncKeyValue `json:"pixel_encoding,omitempty"`
	ScreenInfo    *DesktopInfo `json:"screen_info,omitempty"`
	// server-specified security options
	SecurityTypes []VncKeyValue `json:"security_types,omitempty"`
	Version       *string       `json:"version,omitempty"`
}

func (o *Vnc) GetConnectionFailedReason() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionFailedReason
}

func (o *Vnc) GetDesktopName() *string {
	if o == nil {
		return nil
	}
	return o.DesktopName
}

func (o *Vnc) GetPixelEncoding() *VncKeyValue {
	if o == nil {
		return nil
	}
	return o.PixelEncoding
}

func (o *Vnc) GetScreenInfo() *DesktopInfo {
	if o == nil {
		return nil
	}
	return o.ScreenInfo
}

func (o *Vnc) GetSecurityTypes() []VncKeyValue {
	if o == nil {
		return nil
	}
	return o.SecurityTypes
}

func (o *Vnc) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}
