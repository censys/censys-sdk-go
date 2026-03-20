# RejectType

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.RejectTypeUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.RejectType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `RejectTypeUnknown`           |                               |
| `RejectTypeWrongVersion`      | wrong_version                 |
| `RejectTypeInvalidUsername`   | invalid_username              |
| `RejectTypeWrongUserPw`       | wrong_user_pw                 |
| `RejectTypeWrongServerPw`     | wrong_server_pw               |
| `RejectTypeUsernameInUse`     | username_in_use               |
| `RejectTypeServerFull`        | server_full                   |
| `RejectTypeNoCertificate`     | no_certificate                |
| `RejectTypeAuthenticatorFail` | authenticator_fail            |