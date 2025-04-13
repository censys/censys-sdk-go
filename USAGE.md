<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	censyssdkgointernal "github.com/censys/censys-sdk-go-internal"
	"github.com/censys/censys-sdk-go-internal/models/components"
	"github.com/censys/censys-sdk-go-internal/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := censyssdkgointernal.New(
		censyssdkgointernal.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.GlobalData.Search(ctx, operations.V3GlobaldataSearchQueryRequest{
		OrganizationID: censyssdkgointernal.String("<id>"),
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