# NewStatus

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.NewStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.NewStatus("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `NewStatusActive`     | active                |
| `NewStatusPopulating` | populating            |
| `NewStatusPaused`     | paused                |
| `NewStatusArchived`   | archived              |