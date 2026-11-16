# CollectionPrivacy

Collection visibility. `private` collections are visible to their creator and organization admins. `shared` collections are also visible to all organization members. Only the creator and organization admins can edit a collection.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.CollectionPrivacyPrivate

// Open enum: custom values can be created with a direct type cast
custom := components.CollectionPrivacy("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `CollectionPrivacyPrivate` | private                    |
| `CollectionPrivacyShared`  | shared                     |