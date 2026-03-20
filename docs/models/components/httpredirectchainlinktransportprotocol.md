# HTTPRedirectChainLinkTransportProtocol

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.HTTPRedirectChainLinkTransportProtocolUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.HTTPRedirectChainLinkTransportProtocol("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `HTTPRedirectChainLinkTransportProtocolUnknown` |                                                 |
| `HTTPRedirectChainLinkTransportProtocolTCP`     | tcp                                             |
| `HTTPRedirectChainLinkTransportProtocolUDP`     | udp                                             |
| `HTTPRedirectChainLinkTransportProtocolIcmp`    | icmp                                            |
| `HTTPRedirectChainLinkTransportProtocolQuic`    | quic                                            |