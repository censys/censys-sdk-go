# Source

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.SourceUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.Source("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `SourceUnknown`           |                           |
| `SourceCensys`            | censys                    |
| `SourceRecog`             | recog                     |
| `SourceWappalyzer`        | wappalyzer                |
| `SourceThirdParty`        | third_party               |
| `SourceHTMLMetaExtractor` | html_meta_extractor       |