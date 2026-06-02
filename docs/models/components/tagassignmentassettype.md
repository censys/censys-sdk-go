# TagAssignmentAssetType

The inferred type of the asset.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.TagAssignmentAssetTypeHost

// Open enum: custom values can be created with a direct type cast
custom := components.TagAssignmentAssetType("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `TagAssignmentAssetTypeHost`        | host                                |
| `TagAssignmentAssetTypeCertificate` | certificate                         |
| `TagAssignmentAssetTypeWebProperty` | web_property                        |
| `TagAssignmentAssetTypeUnknown`     | unknown                             |