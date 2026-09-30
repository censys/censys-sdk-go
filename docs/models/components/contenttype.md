# ContentType

The media type of the file to upload. Send it bare: a type carrying parameters, such as a charset, is not accepted.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.ContentTypeApplicationPdf
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ContentTypeApplicationPdf` | application/pdf             |
| `ContentTypeTextHTML`       | text/html                   |
| `ContentTypeTextCsv`        | text/csv                    |
| `ContentTypeTextPlain`      | text/plain                  |
| `ContentTypeImagePng`       | image/png                   |
| `ContentTypeImageJpeg`      | image/jpeg                  |
| `ContentTypeImageWebp`      | image/webp                  |