# ThreatHunting
(*ThreatHunting*)

## Overview

Endpoints related to the Threat Hunting product

### Available Operations

* [GetTrackedScan](#gettrackedscan) - Get tracked scan details
* [GetHostObservationsWithCertificate](#gethostobservationswithcertificate) - Get Host Observations With Certificate
* [CreateTrackedScan](#createtrackedscan) - Create a tracked discovery scan
* [GetTrackedScanThreatHunting](#gettrackedscanthreathunting) - Get tracked scan details
* [ValueCounts](#valuecounts) - CensEye: Retrieve value counts to discover pivots

## GetTrackedScan

Retrieve the current status and results of a tracked scan by its ID.
        This endpoint works for both discovery scans and rescans.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-scans-get" method="get" path="/v3/global/scans/{scan_id}" -->
```go
package main

import(
	"context"
	censyssdkgo "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := censyssdkgo.New(
        censyssdkgo.WithOrganizationID("11111111-2222-3333-4444-555555555555"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.ThreatHunting.GetTrackedScan(ctx, operations.V3GlobaldataScansGetRequest{
        ScanID: "5f39588f-d4c5-48e5-8894-0bb5918c28fa",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeTrackedScan != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.V3GlobaldataScansGetRequest](../../models/operations/v3globaldatascansgetrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.V3GlobaldataScansGetResponse](../../models/operations/v3globaldatascansgetresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## GetHostObservationsWithCertificate

Retrieve historical observations of hosts associated with a certificate fingerprint. Useful for threat hunting, detection engineering, and timeline generation.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-threathunting-get-host-observations-with-certificate" method="get" path="/v3/threat-hunting/certificate/{certificate_id}/observations/hosts" -->
```go
package main

import(
	"context"
	censyssdkgo "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := censyssdkgo.New(
        censyssdkgo.WithOrganizationID("11111111-2222-3333-4444-555555555555"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.ThreatHunting.GetHostObservationsWithCertificate(ctx, operations.V3ThreathuntingGetHostObservationsWithCertificateRequest{
        CertificateID: "55af8a301eb51abdaf7c31bec951638fe5a99d5d92117eca2be493026613fa46",
        StartTime: censyssdkgo.String("2023-01-01T00:00:00Z"),
        EndTime: censyssdkgo.String("2023-12-31T23:59:59Z"),
        Port: censyssdkgo.Int(443),
        Protocol: censyssdkgo.String("TCP"),
        PageSize: censyssdkgo.Int(50),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeHostObservationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.V3ThreathuntingGetHostObservationsWithCertificateRequest](../../models/operations/v3threathuntinggethostobservationswithcertificaterequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |
| `opts`                                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                                   | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.V3ThreathuntingGetHostObservationsWithCertificateResponse](../../models/operations/v3threathuntinggethostobservationswithcertificateresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## CreateTrackedScan

Create a new tracked discovery scan for a specified target. Discovery scans are used to scan new targets that have not been previously identified. The scan will be queued. The response will contain a scan ID that you can use with the [get tracked scan details endpoint](https://docs.censys.com/reference/v3-globaldata-scans-get#/) to monitor its status and results.<br><br>This endpoint is available to organizations that have access to the Threat Hunting module.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-threathunting-scans-discovery" method="post" path="/v3/threat-hunting/scans/discovery" -->
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

    res, err := s.ThreatHunting.CreateTrackedScan(ctx, operations.V3ThreathuntingScansDiscoveryRequest{
        ScansDiscoveryInputBody: components.ScansDiscoveryInputBody{
            Target: components.CreateScansDiscoveryInputBodyTargetTarget2(
                components.Target2{
                    HostnamePort: components.HostnamePort{
                        Hostname: "censys.io",
                        Port: 443,
                    },
                },
            ),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeTrackedScan != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.V3ThreathuntingScansDiscoveryRequest](../../models/operations/v3threathuntingscansdiscoveryrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.V3ThreathuntingScansDiscoveryResponse](../../models/operations/v3threathuntingscansdiscoveryresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## GetTrackedScanThreatHunting

Retrieve the current status and results of a tracked scan by its ID.
        This endpoint works for both discovery scans and rescans.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-threathunting-scans-get" method="get" path="/v3/threat-hunting/scans/{scan_id}" -->
```go
package main

import(
	"context"
	censyssdkgo "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := censyssdkgo.New(
        censyssdkgo.WithOrganizationID("11111111-2222-3333-4444-555555555555"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.ThreatHunting.GetTrackedScanThreatHunting(ctx, operations.V3ThreathuntingScansGetRequest{
        ScanID: "cd62e794-9f12-4c2f-b5b3-153853aaf8d9",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeTrackedScan != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.V3ThreathuntingScansGetRequest](../../models/operations/v3threathuntingscansgetrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.V3ThreathuntingScansGetResponse](../../models/operations/v3threathuntingscansgetresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

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