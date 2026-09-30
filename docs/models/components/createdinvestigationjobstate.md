# CreatedInvestigationJobState

The current state of the investigation. Completed means the investigation succeeded, not that its results are still available to download, and unknown means only that this API is older than the state the service reported.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.CreatedInvestigationJobStateStarted

// Open enum: custom values can be created with a direct type cast
custom := components.CreatedInvestigationJobState("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `CreatedInvestigationJobStateStarted`   | started                                 |
| `CreatedInvestigationJobStateCompleted` | completed                               |
| `CreatedInvestigationJobStateFailed`    | failed                                  |
| `CreatedInvestigationJobStateUnknown`   | unknown                                 |