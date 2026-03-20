# VersionSelected

Certificate version v1(0), v2(1), v3(2).

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.VersionSelectedUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.VersionSelected("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `VersionSelectedUnknown` |                          |
| `VersionSelectedSsLv2`   | ss_lv_2                  |
| `VersionSelectedSsLv3`   | ss_lv_3                  |
| `VersionSelectedTlsv10`  | tlsv1_0                  |
| `VersionSelectedTlsv11`  | tlsv1_1                  |
| `VersionSelectedTlsv12`  | tlsv1_2                  |
| `VersionSelectedTlsv13`  | tlsv1_3                  |
| `VersionSelectedDtlsv10` | dtlsv1_0                 |
| `VersionSelectedDtlsv12` | dtlsv1_2                 |
| `VersionSelectedDtlsv13` | dtlsv1_3                 |