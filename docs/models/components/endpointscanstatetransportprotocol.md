# EndpointScanStateTransportProtocol

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.EndpointScanStateTransportProtocolUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.EndpointScanStateTransportProtocol("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `EndpointScanStateTransportProtocolUnknown` |                                             |
| `EndpointScanStateTransportProtocolTCP`     | tcp                                         |
| `EndpointScanStateTransportProtocolUDP`     | udp                                         |
| `EndpointScanStateTransportProtocolIcmp`    | icmp                                        |
| `EndpointScanStateTransportProtocolQuic`    | quic                                        |