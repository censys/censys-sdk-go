# ServiceIDTransportProtocol

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.ServiceIDTransportProtocolUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.ServiceIDTransportProtocol("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `ServiceIDTransportProtocolUnknown` |                                     |
| `ServiceIDTransportProtocolTCP`     | tcp                                 |
| `ServiceIDTransportProtocolUDP`     | udp                                 |
| `ServiceIDTransportProtocolIcmp`    | icmp                                |
| `ServiceIDTransportProtocolQuic`    | quic                                |