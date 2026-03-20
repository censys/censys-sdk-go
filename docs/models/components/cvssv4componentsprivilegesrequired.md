# CVSSv4ComponentsPrivilegesRequired

Describes the level of privileges or access an attacker must have before successful exploitation. There are three possible values: None (N) – There is no privilege or special access required to conduct the attack, Low (L) – The attacker requires basic, “user” level privileges to leverage the exploit, High (H) – Administrative or similar access privileges are required for successful attack.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.CVSSv4ComponentsPrivilegesRequiredUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.CVSSv4ComponentsPrivilegesRequired("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `CVSSv4ComponentsPrivilegesRequiredUnknown` |                                             |
| `CVSSv4ComponentsPrivilegesRequiredNone`    | none                                        |
| `CVSSv4ComponentsPrivilegesRequiredLow`     | low                                         |
| `CVSSv4ComponentsPrivilegesRequiredHigh`    | high                                        |