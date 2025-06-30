# ThreatHunting
(*ThreatHunting*)

## Overview

Endpoints related to the Threat Hunting product

### Available Operations

* [ValueCounts](#valuecounts) - Value Counts

## ValueCounts

Get counts for specific field-value combinations for threat hunting analysis (requires api-censeye feature flag)

### Example Usage

```go
package main

import(
	"context"
	censyssdkgo "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/models/components"
	"github.com/censys/censys-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := censyssdkgo.New(
        censyssdkgo.WithOrganizationID("<id>"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.ThreatHunting.ValueCounts(ctx, operations.V3ThreathuntingValueCountsRequest{
        SearchValueCountsInputBody: components.SearchValueCountsInputBody{
            AndCountConditions: [][]components.FieldValuePair{
                []components.FieldValuePair{
                    components.FieldValuePair{
                        Field: "<value>",
                        Value: "<value>",
                    },
                },
                []components.FieldValuePair{},
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeValueCountsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.V3ThreathuntingValueCountsRequest](../../models/operations/v3threathuntingvaluecountsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.V3ThreathuntingValueCountsResponse](../../models/operations/v3threathuntingvaluecountsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |