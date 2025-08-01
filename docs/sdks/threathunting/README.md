# ThreatHunting
(*ThreatHunting*)

## Overview

Endpoints related to the Threat Hunting product

### Available Operations

* [ValueCounts](#valuecounts) - CensEye: Retrieve value counts to discover pivots

## ValueCounts

Get counts of web assets for specific field-value pairs and combinations of field-value pairs. This is similar to the [CensEye functionality](https://docs.censys.com/docs/platform-threat-hunting-use-censeye-to-build-detections#/) available in the Platform web UI, but it allows you to define specific fields of interest rather than the [default fields](https://docs.censys.com/docs/platform-threat-hunting-use-censeye-to-build-detections#default-pivot-fields) leveraged by the tool in the UI.<br><br>Each array can only target fields within the same nested object. For example, you can combine `host.services.port=80` and `host.services.protocol=SSH` in the same array, but you cannot combine `host.services.port=80` and `host.location.country=”United States”` in the same array. You can input multiple arrays of objects in each API call.<br><br>To use this endpoint, your organization must have access to the Threat Hunting Module. This endpoint costs 1 credit per count condition (array of objects) included in the API call.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-threathunting-value-counts" method="post" path="/v3/threat-hunting/value-counts" -->
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
        censyssdkgo.WithOrganizationID("11111111-2222-3333-4444-555555555555"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.ThreatHunting.ValueCounts(ctx, operations.V3ThreathuntingValueCountsRequest{
        SearchValueCountsInputBody: components.SearchValueCountsInputBody{
            AndCountConditions: []components.CountCondition{
                components.CountCondition{
                    FieldValuePairs: []components.FieldValuePair{
                        components.FieldValuePair{
                            Field: "host.services.port",
                            Value: "80",
                        },
                    },
                },
                components.CountCondition{
                    FieldValuePairs: []components.FieldValuePair{},
                },
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