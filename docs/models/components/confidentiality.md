# Confidentiality

Refers to the disclosure of sensitive information to authorized and unauthorized users, with the goal being that only authorized users are able to access the target data. Confidentiality has three potential values: High (H) – The attacker has full access to all resources in the impacted system, including highly sensitive information such as encryption keys, Low (L) – The attacker has partial access to information, with no control over what, specifically, they are able to access, None (N) – No data is accessible to unauthorized users as a result of the exploit.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.ConfidentialityUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.Confidentiality("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ConfidentialityUnknown` |                          |
| `ConfidentialityNone`    | none                     |
| `ConfidentialityLow`     | low                      |
| `ConfidentialityHigh`    | high                     |