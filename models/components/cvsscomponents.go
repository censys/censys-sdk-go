// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// AttackComplexity - Indicates conditions beyond the attacker’s control that must exist in order to exploit the vulnerability. The Attack Complexity metric is scored as either Low or High. There are two possible values: Low (L) – There are no specific pre-conditions required for exploitation, High (H) – The attacker must complete some number of preparatory steps in order to get access.
type AttackComplexity string

const (
	AttackComplexityUnknown AttackComplexity = ""
	AttackComplexityLow     AttackComplexity = "low"
	AttackComplexityHigh    AttackComplexity = "high"
)

func (e AttackComplexity) ToPointer() *AttackComplexity {
	return &e
}
func (e *AttackComplexity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "low":
		fallthrough
	case "high":
		*e = AttackComplexity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AttackComplexity: %v", v)
	}
}

// AttackVector - Indicates the level of access required for an attacker to exploit the vulnerability. The Attack Vector metric is scored in one of four levels: Network (N) – Vulnerabilities with this rating are remotely exploitable, from one or more hops away, up to, and including, remote exploitation over the Internet, Adjacent (A) – A vulnerability with this rating requires network adjacency for exploitation. The attack must be launched from the same physical or logical network, Local (L) – Vulnerabilities with this rating are not exploitable over a network, Physical (P) – An attacker must physically interact with the target system.
type AttackVector string

const (
	AttackVectorUnknown  AttackVector = ""
	AttackVectorNetwork  AttackVector = "network"
	AttackVectorAdjacent AttackVector = "adjacent"
	AttackVectorLocal    AttackVector = "local"
	AttackVectorPhysical AttackVector = "physical"
)

func (e AttackVector) ToPointer() *AttackVector {
	return &e
}
func (e *AttackVector) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "network":
		fallthrough
	case "adjacent":
		fallthrough
	case "local":
		fallthrough
	case "physical":
		*e = AttackVector(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AttackVector: %v", v)
	}
}

// Availability - If an attack renders information unavailable, such as when a system crashes or through a DDoS attack, availability is negatively impacted. Availability has three possible values: None (N) – There is no loss of availability, Low (L) – Availability might be intermittently limited, or performance might be negatively impacted, as a result of a successful attack, High (H) – There is a complete loss of availability of the impacted system or information.
type Availability string

const (
	AvailabilityUnknown Availability = ""
	AvailabilityNone    Availability = "none"
	AvailabilityLow     Availability = "low"
	AvailabilityHigh    Availability = "high"
)

func (e Availability) ToPointer() *Availability {
	return &e
}
func (e *Availability) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "none":
		fallthrough
	case "low":
		fallthrough
	case "high":
		*e = Availability(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Availability: %v", v)
	}
}

// Confidentiality - Refers to the disclosure of sensitive information to authorized and unauthorized users, with the goal being that only authorized users are able to access the target data. Confidentiality has three potential values: High (H) – The attacker has full access to all resources in the impacted system, including highly sensitive information such as encryption keys, Low (L) – The attacker has partial access to information, with no control over what, specifically, they are able to access, None (N) – No data is accessible to unauthorized users as a result of the exploit.
type Confidentiality string

const (
	ConfidentialityUnknown Confidentiality = ""
	ConfidentialityNone    Confidentiality = "none"
	ConfidentialityLow     Confidentiality = "low"
	ConfidentialityHigh    Confidentiality = "high"
)

func (e Confidentiality) ToPointer() *Confidentiality {
	return &e
}
func (e *Confidentiality) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "none":
		fallthrough
	case "low":
		fallthrough
	case "high":
		*e = Confidentiality(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Confidentiality: %v", v)
	}
}

// Integrity - Refers to whether the protected information has been tampered with or changed in any way. If there is no way for an attacker to alter the accuracy or completeness of the information, integrity has been maintained. Integrity has three values: None (N) – There is no loss of the integrity of any information, Low (L) – A limited amount of information might be tampered with or modified, but there is no serious impact on the protected system, High (H) – The attacker can modify any/all information on the target system, resulting in a complete loss of integrity.
type Integrity string

const (
	IntegrityUnknown Integrity = ""
	IntegrityNone    Integrity = "none"
	IntegrityLow     Integrity = "low"
	IntegrityHigh    Integrity = "high"
)

func (e Integrity) ToPointer() *Integrity {
	return &e
}
func (e *Integrity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "none":
		fallthrough
	case "low":
		fallthrough
	case "high":
		*e = Integrity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Integrity: %v", v)
	}
}

// PrivilegesRequired - Describes the level of privileges or access an attacker must have before successful exploitation. There are three possible values: None (N) – There is no privilege or special access required to conduct the attack, Low (L) – The attacker requires basic, “user” level privileges to leverage the exploit, High (H) – Administrative or similar access privileges are required for successful attack.
type PrivilegesRequired string

const (
	PrivilegesRequiredUnknown PrivilegesRequired = ""
	PrivilegesRequiredNone    PrivilegesRequired = "none"
	PrivilegesRequiredLow     PrivilegesRequired = "low"
	PrivilegesRequiredHigh    PrivilegesRequired = "high"
)

func (e PrivilegesRequired) ToPointer() *PrivilegesRequired {
	return &e
}
func (e *PrivilegesRequired) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "none":
		fallthrough
	case "low":
		fallthrough
	case "high":
		*e = PrivilegesRequired(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PrivilegesRequired: %v", v)
	}
}

// Scope - Determines whether a vulnerability in one system or component can impact another system or component. If a vulnerability in a vulnerable component can affect a component which is in a different security scope than the vulnerable component, a scope change occurs. Scope has two possible ratings: Changed (C) – An exploited vulnerability can have a carry over impact on another system, Unchanged (U) – The exploited vulnerability is limited in damage to only the local security authority.
type Scope string

const (
	ScopeUnknown   Scope = ""
	ScopeUnchanged Scope = "unchanged"
	ScopeChanged   Scope = "changed"
)

func (e Scope) ToPointer() *Scope {
	return &e
}
func (e *Scope) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "unchanged":
		fallthrough
	case "changed":
		*e = Scope(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Scope: %v", v)
	}
}

// UserInteraction - Describes whether a user, other than the attacker, is required to do anything or participate in exploitation of the vulnerability. User interaction has two possible values: None (N) – No user interaction is required, Required (R) – A user must complete some steps for the exploit to succeed. For example, a user might be required to install some software.
type UserInteraction string

const (
	UserInteractionUnknown  UserInteraction = ""
	UserInteractionNone     UserInteraction = "none"
	UserInteractionRequired UserInteraction = "required"
)

func (e UserInteraction) ToPointer() *UserInteraction {
	return &e
}
func (e *UserInteraction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "none":
		fallthrough
	case "required":
		*e = UserInteraction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserInteraction: %v", v)
	}
}

type CVSSComponents struct {
	// Indicates conditions beyond the attacker’s control that must exist in order to exploit the vulnerability. The Attack Complexity metric is scored as either Low or High. There are two possible values: Low (L) – There are no specific pre-conditions required for exploitation, High (H) – The attacker must complete some number of preparatory steps in order to get access.
	AttackComplexity *AttackComplexity `json:"attack_complexity,omitempty"`
	// Indicates the level of access required for an attacker to exploit the vulnerability. The Attack Vector metric is scored in one of four levels: Network (N) – Vulnerabilities with this rating are remotely exploitable, from one or more hops away, up to, and including, remote exploitation over the Internet, Adjacent (A) – A vulnerability with this rating requires network adjacency for exploitation. The attack must be launched from the same physical or logical network, Local (L) – Vulnerabilities with this rating are not exploitable over a network, Physical (P) – An attacker must physically interact with the target system.
	AttackVector *AttackVector `json:"attack_vector,omitempty"`
	// If an attack renders information unavailable, such as when a system crashes or through a DDoS attack, availability is negatively impacted. Availability has three possible values: None (N) – There is no loss of availability, Low (L) – Availability might be intermittently limited, or performance might be negatively impacted, as a result of a successful attack, High (H) – There is a complete loss of availability of the impacted system or information.
	Availability *Availability `json:"availability,omitempty"`
	// Refers to the disclosure of sensitive information to authorized and unauthorized users, with the goal being that only authorized users are able to access the target data. Confidentiality has three potential values: High (H) – The attacker has full access to all resources in the impacted system, including highly sensitive information such as encryption keys, Low (L) – The attacker has partial access to information, with no control over what, specifically, they are able to access, None (N) – No data is accessible to unauthorized users as a result of the exploit.
	Confidentiality *Confidentiality `json:"confidentiality,omitempty"`
	// Refers to whether the protected information has been tampered with or changed in any way. If there is no way for an attacker to alter the accuracy or completeness of the information, integrity has been maintained. Integrity has three values: None (N) – There is no loss of the integrity of any information, Low (L) – A limited amount of information might be tampered with or modified, but there is no serious impact on the protected system, High (H) – The attacker can modify any/all information on the target system, resulting in a complete loss of integrity.
	Integrity *Integrity `json:"integrity,omitempty"`
	// Describes the level of privileges or access an attacker must have before successful exploitation. There are three possible values: None (N) – There is no privilege or special access required to conduct the attack, Low (L) – The attacker requires basic, “user” level privileges to leverage the exploit, High (H) – Administrative or similar access privileges are required for successful attack.
	PrivilegesRequired *PrivilegesRequired `json:"privileges_required,omitempty"`
	// Determines whether a vulnerability in one system or component can impact another system or component. If a vulnerability in a vulnerable component can affect a component which is in a different security scope than the vulnerable component, a scope change occurs. Scope has two possible ratings: Changed (C) – An exploited vulnerability can have a carry over impact on another system, Unchanged (U) – The exploited vulnerability is limited in damage to only the local security authority.
	Scope *Scope `json:"scope,omitempty"`
	// Describes whether a user, other than the attacker, is required to do anything or participate in exploitation of the vulnerability. User interaction has two possible values: None (N) – No user interaction is required, Required (R) – A user must complete some steps for the exploit to succeed. For example, a user might be required to install some software.
	UserInteraction *UserInteraction `json:"user_interaction,omitempty"`
}

func (o *CVSSComponents) GetAttackComplexity() *AttackComplexity {
	if o == nil {
		return nil
	}
	return o.AttackComplexity
}

func (o *CVSSComponents) GetAttackVector() *AttackVector {
	if o == nil {
		return nil
	}
	return o.AttackVector
}

func (o *CVSSComponents) GetAvailability() *Availability {
	if o == nil {
		return nil
	}
	return o.Availability
}

func (o *CVSSComponents) GetConfidentiality() *Confidentiality {
	if o == nil {
		return nil
	}
	return o.Confidentiality
}

func (o *CVSSComponents) GetIntegrity() *Integrity {
	if o == nil {
		return nil
	}
	return o.Integrity
}

func (o *CVSSComponents) GetPrivilegesRequired() *PrivilegesRequired {
	if o == nil {
		return nil
	}
	return o.PrivilegesRequired
}

func (o *CVSSComponents) GetScope() *Scope {
	if o == nil {
		return nil
	}
	return o.Scope
}

func (o *CVSSComponents) GetUserInteraction() *UserInteraction {
	if o == nil {
		return nil
	}
	return o.UserInteraction
}
