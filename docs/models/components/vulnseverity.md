# VulnSeverity

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.VulnSeverityUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.VulnSeverity("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `VulnSeverityUnknown`  |                        |
| `VulnSeverityLow`      | low                    |
| `VulnSeverityMedium`   | medium                 |
| `VulnSeverityHigh`     | high                   |
| `VulnSeverityCritical` | critical               |