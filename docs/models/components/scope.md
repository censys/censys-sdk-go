# Scope

Determines whether a vulnerability in one system or component can impact another system or component. If a vulnerability in a vulnerable component can affect a component which is in a different security scope than the vulnerable component, a scope change occurs. Scope has two possible ratings: Changed (C) – An exploited vulnerability can have a carry over impact on another system, Unchanged (U) – The exploited vulnerability is limited in damage to only the local security authority.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.ScopeUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.Scope("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `ScopeUnknown`   |                  |
| `ScopeUnchanged` | unchanged        |
| `ScopeChanged`   | changed          |