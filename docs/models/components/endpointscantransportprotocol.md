# EndpointScanTransportProtocol

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.EndpointScanTransportProtocolUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.EndpointScanTransportProtocol("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `EndpointScanTransportProtocolUnknown` |                                        |
| `EndpointScanTransportProtocolTCP`     | tcp                                    |
| `EndpointScanTransportProtocolUDP`     | udp                                    |
| `EndpointScanTransportProtocolIcmp`    | icmp                                   |
| `EndpointScanTransportProtocolQuic`    | quic                                   |