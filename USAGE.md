<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	censyssdkgo "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/models/components"
	"github.com/censys/censys-sdk-go/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := censyssdkgo.New(
		censyssdkgo.WithOrganizationID("11111111-2222-3333-4444-555555555555"),
		censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.GlobalData.Search(ctx, operations.V3GlobaldataSearchQueryRequest{
		SearchQueryInputBody: components.SearchQueryInputBody{
			Fields: []string{
				"host.ip",
			},
			PageSize: censyssdkgo.Pointer[int64](1),
			Query:    "host.services: (protocol=SSH and not port: 22)",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseEnvelopeSearchQueryResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->