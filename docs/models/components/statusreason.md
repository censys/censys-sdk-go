# StatusReason

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.StatusReasonUnspecified

// Open enum: custom values can be created with a direct type cast
custom := components.StatusReason("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `StatusReasonUnspecified`      | unspecified                    |
| `StatusReasonNotEnoughCredits` | not_enough_credits             |
| `StatusReasonNotEntitled`      | not_entitled                   |
| `StatusReasonTooManyAssets`    | too_many_assets                |
| `StatusReasonManual`           | manual                         |
| `StatusReasonQueryChanged`     | query_changed                  |
| `StatusReasonInitial`          | initial                        |