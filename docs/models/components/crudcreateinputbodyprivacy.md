# CrudCreateInputBodyPrivacy

Collection visibility. `private` collections are visible to their creator and organization admins. `shared` collections are also visible to all organization members. Only the creator and organization admins can edit a collection.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.CrudCreateInputBodyPrivacyPrivate
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `CrudCreateInputBodyPrivacyPrivate` | private                             |
| `CrudCreateInputBodyPrivacyShared`  | shared                              |