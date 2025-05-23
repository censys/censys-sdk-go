// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type SSHServerHostKey struct {
	CertkeyPublicKey  *string                  `json:"certkey_public_key,omitempty"`
	DsaPublicKey      *DSACryptographicKey     `json:"dsa_public_key,omitempty"`
	EcdsaPublicKey    *ECDSACryptographicKey   `json:"ecdsa_public_key,omitempty"`
	Ed25519PublicKey  *Ed25519CryptographicKey `json:"ed25519_public_key,omitempty"`
	FingerprintSha256 *string                  `json:"fingerprint_sha256,omitempty"`
	RsaPublicKey      *RSACryptographicKey     `json:"rsa_public_key,omitempty"`
}

func (o *SSHServerHostKey) GetCertkeyPublicKey() *string {
	if o == nil {
		return nil
	}
	return o.CertkeyPublicKey
}

func (o *SSHServerHostKey) GetDsaPublicKey() *DSACryptographicKey {
	if o == nil {
		return nil
	}
	return o.DsaPublicKey
}

func (o *SSHServerHostKey) GetEcdsaPublicKey() *ECDSACryptographicKey {
	if o == nil {
		return nil
	}
	return o.EcdsaPublicKey
}

func (o *SSHServerHostKey) GetEd25519PublicKey() *Ed25519CryptographicKey {
	if o == nil {
		return nil
	}
	return o.Ed25519PublicKey
}

func (o *SSHServerHostKey) GetFingerprintSha256() *string {
	if o == nil {
		return nil
	}
	return o.FingerprintSha256
}

func (o *SSHServerHostKey) GetRsaPublicKey() *RSACryptographicKey {
	if o == nil {
		return nil
	}
	return o.RsaPublicKey
}
