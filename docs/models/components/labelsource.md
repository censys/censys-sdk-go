# LabelSource

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.LabelSourceUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.LabelSource("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `LabelSourceUnknown`           |                                |
| `LabelSourceCensys`            | censys                         |
| `LabelSourceRecog`             | recog                          |
| `LabelSourceWappalyzer`        | wappalyzer                     |
| `LabelSourceThirdParty`        | third_party                    |
| `LabelSourceHTMLMetaExtractor` | html_meta_extractor            |