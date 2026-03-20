# ServerType

An enumerated value indicating the behavior of the server. An AUTHORITATIVE server fulfills requests for domain names it controls, which are not listed by the server. FORWARDING and RECURSIVE_RESOLVER servers fulfill requests indirectly for domain names they do not control. A RECURSIVE_RESOLVER will query ip.parrotdns.com itself, resulting in its own IP address being present in the dns.answers.response field.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.ServerTypeUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.ServerType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `ServerTypeUnknown`           |                               |
| `ServerTypeRecursiveResolver` | recursive_resolver            |
| `ServerTypeAuthoritative`     | authoritative                 |
| `ServerTypeForwarding`        | forwarding                    |
| `ServerTypeRedirecting`       | redirecting                   |