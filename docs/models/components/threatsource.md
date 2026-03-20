# ThreatSource

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.ThreatSourceUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.ThreatSource("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ThreatSourceUnknown`           |                                 |
| `ThreatSourceCensys`            | censys                          |
| `ThreatSourceRecog`             | recog                           |
| `ThreatSourceWappalyzer`        | wappalyzer                      |
| `ThreatSourceThirdParty`        | third_party                     |
| `ThreatSourceHTMLMetaExtractor` | html_meta_extractor             |