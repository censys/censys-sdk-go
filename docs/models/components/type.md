# Type

The certificate's type. Options include root, intermediate, or leaf.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.TypeUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.Type("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `TypeUnknown`      |                    |
| `TypeRoot`         | root               |
| `TypeIntermediate` | intermediate       |
| `TypeLeaf`         | leaf               |