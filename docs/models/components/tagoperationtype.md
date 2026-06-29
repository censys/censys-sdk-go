# TagOperationType

Whether the operation creates or deletes tag assignments.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.TagOperationTypeBulkCreate

// Open enum: custom values can be created with a direct type cast
custom := components.TagOperationType("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `TagOperationTypeBulkCreate` | bulk_create                  |
| `TagOperationTypeBulkDelete` | bulk_delete                  |