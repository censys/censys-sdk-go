# ServiceTransportProtocol

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.ServiceTransportProtocolUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.ServiceTransportProtocol("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ServiceTransportProtocolUnknown` |                                   |
| `ServiceTransportProtocolTCP`     | tcp                               |
| `ServiceTransportProtocolUDP`     | udp                               |
| `ServiceTransportProtocolIcmp`    | icmp                              |
| `ServiceTransportProtocolQuic`    | quic                              |