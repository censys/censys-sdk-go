# RiskSource1

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.RiskSource1Unknown

// Open enum: custom values can be created with a direct type cast
custom := components.RiskSource1("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `RiskSource1Unknown`           |                                |
| `RiskSource1Censys`            | censys                         |
| `RiskSource1Recog`             | recog                          |
| `RiskSource1Wappalyzer`        | wappalyzer                     |
| `RiskSource1ThirdParty`        | third_party                    |
| `RiskSource1HTMLMetaExtractor` | html_meta_extractor            |