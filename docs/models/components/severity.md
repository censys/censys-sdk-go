# Severity

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.SeverityUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.Severity("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `SeverityUnknown`  |                    |
| `SeverityLow`      | low                |
| `SeverityMedium`   | medium             |
| `SeverityHigh`     | high               |
| `SeverityCritical` | critical           |