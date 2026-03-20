# CollectionStatus

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.CollectionStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := components.CollectionStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `CollectionStatusUnspecified` | unspecified                   |
| `CollectionStatusPopulating`  | populating                    |
| `CollectionStatusActive`      | active                        |
| `CollectionStatusPaused`      | paused                        |
| `CollectionStatusArchived`    | archived                      |