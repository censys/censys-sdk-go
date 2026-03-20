# ProviderUrgency

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.ProviderUrgencyUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.ProviderUrgency("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ProviderUrgencyUnknown` |                          |
| `ProviderUrgencyClear`   | clear                    |
| `ProviderUrgencyGreen`   | green                    |
| `ProviderUrgencyAmber`   | amber                    |
| `ProviderUrgencyRed`     | red                      |