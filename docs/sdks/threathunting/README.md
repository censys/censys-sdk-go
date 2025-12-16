# ThreatHunting

## Overview

Endpoints related to the Threat Hunting product

### Available Operations

* [GetHostObservationsWithCertificate](#gethostobservationswithcertificate) - Get host history for a certificate
* [CreateTrackedScan](#createtrackedscan) - Live Discovery: Initiate a new scan
* [GetTrackedScanThreatHunting](#gettrackedscanthreathunting) - Get scan status
* [ListThreats](#listthreats) - List active threats
* [ValueCounts](#valuecounts) - CensEye: Retrieve value counts to discover pivots

## GetHostObservationsWithCertificate

Retrieve the historical observations of hosts associated with a certificate. This is useful for threat hunting, detection engineering, and timeline generation. Certificate history is also visible to Threat Hunting users in the Platform UI on the [certificate timeline](https://docs.censys.com/docs/platform-threat-hunting-use-cert-history-to-build-better-detections#/).<br><br>You can define a specific time frame of interest. If you do not specify a time frame, this endpoint will search the historical dataset that is available to your account. You may also filter results by port and transport protocol.<br><br>This endpoint is available to organizations that have access to the Threat Hunting module. It costs 5 credits per page of results.

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
        StartTime: censyssdkgo.Pointer("2023-01-01T00:00:00Z"),
        EndTime: censyssdkgo.Pointer("2023-12-31T23:59:59Z"),
        Port: censyssdkgo.Pointer[int](443),
        Protocol: censyssdkgo.Pointer("TCP"),
        PageSize: censyssdkgo.Pointer[int](50),
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

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## CreateTrackedScan

Initiate a scan to look for a currently unobserved service at a specific IP and port (`ip:port`) or hostname and port (`hostname:port`). This is equivalent to the [Live Discovery](https://docs.censys.com/docs/platform-threat-hunting-use-live-scan-and-rescan-to-validate-infrastructure#/) feature available in the UI, but you can also target web properties in addition to hosts.<br><br>The scan may take several minutes to complete. The response will contain a scan ID that you can use to [monitor the scan's status](https://docs.censys.com/reference/v3-threathunting-scans-get#/). After the scan completes, perform a lookup on the target asset to retrieve detailed scan information.<br><br>This endpoint is available to organizations that have access to the Threat Hunting module. It costs 15 credits to execute this endpoint.

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

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetTrackedScanThreatHunting

Retrieve the current status of a scan by its ID. This endpoint works for both [Live Discovery scans](https://docs.censys.com/reference/v3-threathunting-scans-discovery#/) and [Live Rescans](https://docs.censys.com/reference/v3-globaldata-scans-rescan#/).<br><br>If the scan was successful, perform a lookup on the target asset to retrieve detailed scan information.<br><br>This endpoint is available to all Enterprise customers. This endpoint does not cost any credits to execute.

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

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404                      | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## ListThreats

Retrieve a list of active threats observed by Censys by aggregating threat IDs across hosts and web properties. Threats are active if their fingerprint has been identified on hosts or web properties by Censys scans. This information is also available on the [Explore Threats page in the Platform web UI](https://platform.censys.io/threats).

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-threathunting-threats-list" method="get" path="/v3/threat-hunting/threats" -->
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

    res, err := s.ThreatHunting.ListThreats(ctx, operations.V3ThreathuntingThreatsListRequest{
        Query: censyssdkgo.Pointer("*"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeThreatsListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.V3ThreathuntingThreatsListRequest](../../models/operations/v3threathuntingthreatslistrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.V3ThreathuntingThreatsListResponse](../../models/operations/v3threathuntingthreatslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 422                      | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

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

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |