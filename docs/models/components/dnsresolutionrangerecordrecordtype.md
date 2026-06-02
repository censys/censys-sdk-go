# DNSResolutionRangeRecordRecordType

The record type. Either A, AAAA, MX, NS, SOA, or TXT.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.DNSResolutionRangeRecordRecordTypeA

// Open enum: custom values can be created with a direct type cast
custom := components.DNSResolutionRangeRecordRecordType("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `DNSResolutionRangeRecordRecordTypeA`    | A                                        |
| `DNSResolutionRangeRecordRecordTypeAaaa` | AAAA                                     |
| `DNSResolutionRangeRecordRecordTypeMx`   | MX                                       |
| `DNSResolutionRangeRecordRecordTypeNs`   | NS                                       |
| `DNSResolutionRangeRecordRecordTypeSoa`  | SOA                                      |
| `DNSResolutionRangeRecordRecordTypeTxt`  | TXT                                      |