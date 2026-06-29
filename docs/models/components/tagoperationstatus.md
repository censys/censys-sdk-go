# TagOperationStatus

The current status of the operation.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.TagOperationStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.TagOperationStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `TagOperationStatusPending`      | pending                          |
| `TagOperationStatusRunning`      | running                          |
| `TagOperationStatusSucceeded`    | succeeded                        |
| `TagOperationStatusLimitReached` | limit_reached                    |
| `TagOperationStatusFailed`       | failed                           |
| `TagOperationStatusCancelled`    | cancelled                        |