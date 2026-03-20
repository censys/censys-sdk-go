# VulnSource

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.VulnSourceUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.VulnSource("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `VulnSourceUnknown`           |                               |
| `VulnSourceCensys`            | censys                        |
| `VulnSourceRecog`             | recog                         |
| `VulnSourceWappalyzer`        | wappalyzer                    |
| `VulnSourceThirdParty`        | third_party                   |
| `VulnSourceHTMLMetaExtractor` | html_meta_extractor           |