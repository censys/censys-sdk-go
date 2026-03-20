# KEVSource

The source checked to determine whether the CVE is in the KEV catalog.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.KEVSourceUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.KEVSource("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `KEVSourceUnknown`    |                       |
| `KEVSourceCisa`       | cisa                  |
| `KEVSourceThirdParty` | third_party           |