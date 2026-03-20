# DNSResourceRecordType

An enumerated field indicating what type of data is in the "services.dns.additionals.response" field. For example, "A" signifies that the value in "services.dns.additionals.response" is an IPv4 address for the FQDN in "services.dns.additionals.name".

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.DNSResourceRecordTypeUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.DNSResourceRecordType("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `DNSResourceRecordTypeUnknown` |                                |
| `DNSResourceRecordTypeA`       | a                              |
| `DNSResourceRecordTypeTxt`     | txt                            |
| `DNSResourceRecordTypeNs`      | ns                             |