# CVSSv4ComponentsAttackComplexity

Indicates conditions beyond the attacker’s control that must exist in order to exploit the vulnerability. The Attack Complexity metric is scored as either Low or High. There are two possible values: Low (L) – There are no specific pre-conditions required for exploitation, High (H) – The attacker must complete some number of preparatory steps in order to get access.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.CVSSv4ComponentsAttackComplexityUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.CVSSv4ComponentsAttackComplexity("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `CVSSv4ComponentsAttackComplexityUnknown` |                                           |
| `CVSSv4ComponentsAttackComplexityLow`     | low                                       |
| `CVSSv4ComponentsAttackComplexityHigh`    | high                                      |