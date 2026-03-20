# Granularity

The granularity of the report.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.GranularityDaily

// Open enum: custom values can be created with a direct type cast
custom := components.Granularity("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `GranularityDaily`   | daily                |
| `GranularityMonthly` | monthly              |