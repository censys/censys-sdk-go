# RiskSource

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.RiskSourceUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.RiskSource("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `RiskSourceUnknown` |                     |
| `RiskSourceCensys`  | censys              |
| `RiskSourceCve`     | cve                 |