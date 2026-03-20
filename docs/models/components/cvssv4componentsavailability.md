# CVSSv4ComponentsAvailability

If an attack renders information unavailable, such as when a system crashes or through a DDoS attack, availability is negatively impacted. Availability has three possible values: None (N) – There is no loss of availability, Low (L) – Availability might be intermittently limited, or performance might be negatively impacted, as a result of a successful attack, High (H) – There is a complete loss of availability of the impacted system or information.

## Example Usage

```go
import (
	"github.com/censys/censys-sdk-go/models/components"
)

value := components.CVSSv4ComponentsAvailabilityUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.CVSSv4ComponentsAvailability("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `CVSSv4ComponentsAvailabilityUnknown` |                                       |
| `CVSSv4ComponentsAvailabilityNone`    | none                                  |
| `CVSSv4ComponentsAvailabilityLow`     | low                                   |
| `CVSSv4ComponentsAvailabilityHigh`    | high                                  |