# State

Current state of the job.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.StateStarted

// Open enum: custom values can be created with a direct type cast
custom := components.State("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `StateStarted`   | started          |
| `StateCompleted` | completed        |
| `StateFailed`    | failed           |
| `StateUnknown`   | unknown          |