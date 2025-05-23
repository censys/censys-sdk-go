// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CertificateExtensions struct {
	AuthorityInfoAccess *AuthorityInfoAccess `json:"authority_info_access,omitempty"`
	// A key identifier, usually a digest of the DER-encoded SubjectPublicKeyInfo.
	AuthorityKeyID     *string             `json:"authority_key_id,omitempty"`
	BasicConstraints   *BasicConstraints   `json:"basic_constraints,omitempty"`
	CabfOrganizationID *CabfOrganizationID `json:"cabf_organization_id,omitempty"`
	// The parsed id-ce-certificatePolicies extension (OID: 2.5.29.32).
	CertificatePolicies []CertificatePolicy `json:"certificate_policies,omitempty"`
	// The parsed id-ce-cRLDistributionPoints extension (OID: 2.5.29.31). Contents are a list of distributionPoint URLs; other distributionPoint types are omitted).
	CrlDistributionPoints []string `json:"crl_distribution_points,omitempty"`
	// Whether the certificate possesses the pre-certificate "poison" extension (OID: 1.3.6.1.4.1.11129.2.4.3).
	CtPoison                    *bool                        `json:"ct_poison,omitempty"`
	ExtendedKeyUsage            *ExtendedKeyUsage            `json:"extended_key_usage,omitempty"`
	IssuerAltName               *GeneralNames                `json:"issuer_alt_name,omitempty"`
	KeyUsage                    *KeyUsage                    `json:"key_usage,omitempty"`
	NameConstraints             *NameConstraints             `json:"name_constraints,omitempty"`
	QcStatements                *QcStatements                `json:"qc_statements,omitempty"`
	SignedCertificateTimestamps []SignedCertificateTimestamp `json:"signed_certificate_timestamps,omitempty"`
	SubjectAltName              *GeneralNames                `json:"subject_alt_name,omitempty"`
	// A key identifier, usually a digest of the DER-encoded SubjectPublicKeyInfo..
	SubjectKeyID          *string                `json:"subject_key_id,omitempty"`
	TorServiceDescriptors []TorServiceDescriptor `json:"tor_service_descriptors,omitempty"`
}

func (o *CertificateExtensions) GetAuthorityInfoAccess() *AuthorityInfoAccess {
	if o == nil {
		return nil
	}
	return o.AuthorityInfoAccess
}

func (o *CertificateExtensions) GetAuthorityKeyID() *string {
	if o == nil {
		return nil
	}
	return o.AuthorityKeyID
}

func (o *CertificateExtensions) GetBasicConstraints() *BasicConstraints {
	if o == nil {
		return nil
	}
	return o.BasicConstraints
}

func (o *CertificateExtensions) GetCabfOrganizationID() *CabfOrganizationID {
	if o == nil {
		return nil
	}
	return o.CabfOrganizationID
}

func (o *CertificateExtensions) GetCertificatePolicies() []CertificatePolicy {
	if o == nil {
		return nil
	}
	return o.CertificatePolicies
}

func (o *CertificateExtensions) GetCrlDistributionPoints() []string {
	if o == nil {
		return nil
	}
	return o.CrlDistributionPoints
}

func (o *CertificateExtensions) GetCtPoison() *bool {
	if o == nil {
		return nil
	}
	return o.CtPoison
}

func (o *CertificateExtensions) GetExtendedKeyUsage() *ExtendedKeyUsage {
	if o == nil {
		return nil
	}
	return o.ExtendedKeyUsage
}

func (o *CertificateExtensions) GetIssuerAltName() *GeneralNames {
	if o == nil {
		return nil
	}
	return o.IssuerAltName
}

func (o *CertificateExtensions) GetKeyUsage() *KeyUsage {
	if o == nil {
		return nil
	}
	return o.KeyUsage
}

func (o *CertificateExtensions) GetNameConstraints() *NameConstraints {
	if o == nil {
		return nil
	}
	return o.NameConstraints
}

func (o *CertificateExtensions) GetQcStatements() *QcStatements {
	if o == nil {
		return nil
	}
	return o.QcStatements
}

func (o *CertificateExtensions) GetSignedCertificateTimestamps() []SignedCertificateTimestamp {
	if o == nil {
		return nil
	}
	return o.SignedCertificateTimestamps
}

func (o *CertificateExtensions) GetSubjectAltName() *GeneralNames {
	if o == nil {
		return nil
	}
	return o.SubjectAltName
}

func (o *CertificateExtensions) GetSubjectKeyID() *string {
	if o == nil {
		return nil
	}
	return o.SubjectKeyID
}

func (o *CertificateExtensions) GetTorServiceDescriptors() []TorServiceDescriptor {
	if o == nil {
		return nil
	}
	return o.TorServiceDescriptors
}
