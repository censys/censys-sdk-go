# UserInteraction

Describes whether a user, other than the attacker, is required to do anything or participate in exploitation of the vulnerability. User interaction has two possible values: None (N) – No user interaction is required, Required (R) – A user must complete some steps for the exploit to succeed. For example, a user might be required to install some software.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.UserInteractionUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.UserInteraction("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `UserInteractionUnknown`  |                           |
| `UserInteractionNone`     | none                      |
| `UserInteractionRequired` | required                  |