# Safety

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.SafetyUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.Safety("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `SafetyUnknown`    |                    |
| `SafetyNegligible` | negligible         |
| `SafetyPresent`    | present            |