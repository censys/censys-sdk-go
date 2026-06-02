# DNSResolutionRecordRecordType

The record type. Either A, AAAA, MX, NS, SOA, or TXT.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.DNSResolutionRecordRecordTypeA

// Open enum: custom values can be created with a direct type cast
custom := components.DNSResolutionRecordRecordType("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `DNSResolutionRecordRecordTypeA`    | A                                   |
| `DNSResolutionRecordRecordTypeAaaa` | AAAA                                |
| `DNSResolutionRecordRecordTypeMx`   | MX                                  |
| `DNSResolutionRecordRecordTypeNs`   | NS                                  |
| `DNSResolutionRecordRecordTypeSoa`  | SOA                                 |
| `DNSResolutionRecordRecordTypeTxt`  | TXT                                 |