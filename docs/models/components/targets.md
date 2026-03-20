# Targets

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.TargetsHost

// Open enum: custom values can be created with a direct type cast
custom := components.Targets("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `TargetsHost`    | host             |
| `TargetsWeb`     | web              |
| `TargetsCert`    | cert             |
| `TargetsUnknown` | unknown          |