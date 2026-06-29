# Status

Filter by operation status.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/operations"
)

value := operations.StatusPending
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `StatusPending`      | pending              |
| `StatusRunning`      | running              |
| `StatusSucceeded`    | succeeded            |
| `StatusLimitReached` | limit_reached        |
| `StatusFailed`       | failed               |
| `StatusCancelled`    | cancelled            |