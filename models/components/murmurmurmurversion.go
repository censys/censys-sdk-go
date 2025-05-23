// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type MurmurMurmurVersion struct {
	Build         *int    `json:"build,omitempty"`
	Major         *int    `json:"major,omitempty"`
	Minor         *int    `json:"minor,omitempty"`
	Os            *string `json:"os,omitempty"`
	OsVersion     *string `json:"os_version,omitempty"`
	VersionString *string `json:"version_string,omitempty"`
}

func (o *MurmurMurmurVersion) GetBuild() *int {
	if o == nil {
		return nil
	}
	return o.Build
}

func (o *MurmurMurmurVersion) GetMajor() *int {
	if o == nil {
		return nil
	}
	return o.Major
}

func (o *MurmurMurmurVersion) GetMinor() *int {
	if o == nil {
		return nil
	}
	return o.Minor
}

func (o *MurmurMurmurVersion) GetOs() *string {
	if o == nil {
		return nil
	}
	return o.Os
}

func (o *MurmurMurmurVersion) GetOsVersion() *string {
	if o == nil {
		return nil
	}
	return o.OsVersion
}

func (o *MurmurMurmurVersion) GetVersionString() *string {
	if o == nil {
		return nil
	}
	return o.VersionString
}
