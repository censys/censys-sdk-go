# ParseStatus

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.ParseStatusUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.ParseStatus("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `ParseStatusUnknown`   |                        |
| `ParseStatusSuccess`   | success                |
| `ParseStatusFail`      | fail                   |
| `ParseStatusCorrupted` | corrupted              |