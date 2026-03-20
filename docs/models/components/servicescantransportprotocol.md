# ServiceScanTransportProtocol

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.ServiceScanTransportProtocolUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.ServiceScanTransportProtocol("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `ServiceScanTransportProtocolUnknown` |                                       |
| `ServiceScanTransportProtocolTCP`     | tcp                                   |
| `ServiceScanTransportProtocolUDP`     | udp                                   |
| `ServiceScanTransportProtocolIcmp`    | icmp                                  |
| `ServiceScanTransportProtocolQuic`    | quic                                  |