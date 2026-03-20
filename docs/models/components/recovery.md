# Recovery

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.RecoveryUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.Recovery("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `RecoveryUnknown`       |                         |
| `RecoveryAutomatic`     | automatic               |
| `RecoveryUser`          | user                    |
| `RecoveryIrrecoverable` | irrecoverable           |