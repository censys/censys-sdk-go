# AssetType

The inferred type of the asset.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.AssetTypeHost

// Open enum: custom values can be created with a direct type cast
custom := components.AssetType("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `AssetTypeHost`        | host                   |
| `AssetTypeCertificate` | certificate            |
| `AssetTypeWebProperty` | web_property           |
| `AssetTypeUnknown`     | unknown                |