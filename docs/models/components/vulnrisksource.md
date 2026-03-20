# VulnRiskSource

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.VulnRiskSourceUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.VulnRiskSource("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `VulnRiskSourceUnknown` |                         |
| `VulnRiskSourceCensys`  | censys                  |
| `VulnRiskSourceCve`     | cve                     |