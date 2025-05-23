// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// Type - The certificate's type. Options include root, intermediate, or leaf.
type Type string

const (
	TypeUnknown      Type = ""
	TypeRoot         Type = "root"
	TypeIntermediate Type = "intermediate"
	TypeLeaf         Type = "leaf"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "root":
		fallthrough
	case "intermediate":
		fallthrough
	case "leaf":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type RootStore struct {
	// A path of trusted signing certificates up to a root certificate present in a root store, represented as an ordered list of SHA-256 fingerprints.
	Chains []RootStoreChain `json:"chains,omitempty"`
	// Whether the certificate has ever been considered valid by the root store.
	EverValid *bool `json:"ever_valid,omitempty"`
	// Whether there ever existed a trusted path of signing certificates from a certificate present in the root certificate store.
	HadTrustedPath *bool `json:"had_trusted_path,omitempty"`
	// Whether there currently exists a trusted path of signing certificates from a certificate present in the root certificate store.
	HasTrustedPath *bool `json:"has_trusted_path,omitempty"`
	// Whether the certificate is in the revocation set (e.g. OneCRL) associated with the root store.
	InRevocationSet *bool `json:"in_revocation_set,omitempty"`
	// Whether the certificate is currently considered valid by the root store: a summary of the trust path, revoked, blocklisted/allowlisted, and expired fields.
	IsValid *bool `json:"is_valid,omitempty"`
	// The SHA-256 fingerprints of the certificate's immediate parents in its trust path(s).
	Parents []string `json:"parents,omitempty"`
	// The certificate's type. Options include root, intermediate, or leaf.
	Type *Type `json:"type,omitempty"`
}

func (o *RootStore) GetChains() []RootStoreChain {
	if o == nil {
		return nil
	}
	return o.Chains
}

func (o *RootStore) GetEverValid() *bool {
	if o == nil {
		return nil
	}
	return o.EverValid
}

func (o *RootStore) GetHadTrustedPath() *bool {
	if o == nil {
		return nil
	}
	return o.HadTrustedPath
}

func (o *RootStore) GetHasTrustedPath() *bool {
	if o == nil {
		return nil
	}
	return o.HasTrustedPath
}

func (o *RootStore) GetInRevocationSet() *bool {
	if o == nil {
		return nil
	}
	return o.InRevocationSet
}

func (o *RootStore) GetIsValid() *bool {
	if o == nil {
		return nil
	}
	return o.IsValid
}

func (o *RootStore) GetParents() []string {
	if o == nil {
		return nil
	}
	return o.Parents
}

func (o *RootStore) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}
