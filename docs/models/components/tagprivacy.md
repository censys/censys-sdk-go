# TagPrivacy

Tag visibility and management settings. `private` tags are only visible to and editable by organization admins. `shared` tags are visible to and editable by all organization members.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.TagPrivacyPrivate

// Open enum: custom values can be created with a direct type cast
custom := components.TagPrivacy("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `TagPrivacyPrivate` | private             |
| `TagPrivacyShared`  | shared              |