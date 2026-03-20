# Reason

An enumerated value indicating the issuer-supplied reason for the revocation.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.ReasonUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.Reason("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `ReasonUnknown`              |                              |
| `ReasonUnspecified`          | unspecified                  |
| `ReasonKeyCompromise`        | key_compromise               |
| `ReasonCaCompromise`         | ca_compromise                |
| `ReasonAffiliationChanged`   | affiliation_changed          |
| `ReasonSuperseded`           | superseded                   |
| `ReasonCessationOfOperation` | cessation_of_operation       |
| `ReasonCertificateHold`      | certificate_hold             |
| `ReasonRemoveFromCrl`        | remove_from_crl              |
| `ReasonPrivilegeWithdrawn`   | privilege_withdrawn          |
| `ReasonAaCompromise`         | aa_compromise                |