# Version

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.VersionUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.Version("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `VersionUnknown` |                  |
| `VersionSsLv2`   | ss_lv_2          |
| `VersionSsLv3`   | ss_lv_3          |
| `VersionTlsv10`  | tlsv1_0          |
| `VersionTlsv11`  | tlsv1_1          |
| `VersionTlsv12`  | tlsv1_2          |
| `VersionTlsv13`  | tlsv1_3          |
| `VersionDtlsv10` | dtlsv1_0         |
| `VersionDtlsv12` | dtlsv1_2         |
| `VersionDtlsv13` | dtlsv1_3         |