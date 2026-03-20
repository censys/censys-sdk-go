# ScoreLevel

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.ScoreLevelUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.ScoreLevel("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `ScoreLevelUnknown`    |                        |
| `ScoreLevelBenign`     | benign                 |
| `ScoreLevelLowRisk`    | low_risk               |
| `ScoreLevelMediumRisk` | medium_risk            |
| `ScoreLevelHighRisk`   | high_risk              |
| `ScoreLevelMalicious`  | malicious              |