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
		censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.GlobalData.Search(ctx, operations.V3GlobaldataSearchQueryRequest{
		OrganizationID: censyssdkgo.String("<id>"),
		SearchQueryInputBody: components.SearchQueryInputBody{
			Query: "<value>",
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