# TrackedScanTaskStatus

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.TrackedScanTaskStatusUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.TrackedScanTaskStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `TrackedScanTaskStatusUnknown`   |                                  |
| `TrackedScanTaskStatusScanning`  | scanning                         |
| `TrackedScanTaskStatusScanned`   | scanned                          |
| `TrackedScanTaskStatusRejected`  | rejected                         |
| `TrackedScanTaskStatusTimedOut`  | timed_out                        |
| `TrackedScanTaskStatusCompleted` | completed                        |
| `TrackedScanTaskStatusIgnored`   | ignored                          |