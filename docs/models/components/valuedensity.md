# ValueDensity

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.ValueDensityUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.ValueDensity("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ValueDensityUnknown`      |                            |
| `ValueDensityDiffuse`      | diffuse                    |
| `ValueDensityConcentrated` | concentrated               |