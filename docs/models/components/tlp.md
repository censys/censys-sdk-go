# Tlp

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.TlpUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.Tlp("custom_value")
```


## Values

| Name         | Value        |
| ------------ | ------------ |
| `TlpUnknown` |              |
| `TlpGreen`   | green        |
| `TlpAmber`   | amber        |
| `TlpRed`     | red          |
| `TlpWhite`   | white        |