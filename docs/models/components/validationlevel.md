# ValidationLevel

The extent to which the certificate's issuer validated the identity of the entity requesting the certificate. Options include Domain validated (DV), Organization Validated (OV), or Extended Validation (EV).

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.ValidationLevelUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.ValidationLevel("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ValidationLevelUnknown` |                          |
| `ValidationLevelDv`      | dv                       |
| `ValidationLevelOv`      | ov                       |
| `ValidationLevelEv`      | ev                       |