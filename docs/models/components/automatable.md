# Automatable

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.AutomatableUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.Automatable("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `AutomatableUnknown` |                      |
| `AutomatableNo`      | no                   |
| `AutomatableYes`     | yes                  |