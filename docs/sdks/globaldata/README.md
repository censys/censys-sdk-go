# GlobalData
(*GlobalData*)

## Overview

Endpoints related to the Global Data product

### Available Operations

* [GetCertificates](#getcertificates) - Get multiple certificates
* [GetCertificate](#getcertificate) - Get a certificate
* [GetHostObservationsWithCertificate](#gethostobservationswithcertificate) - Get Host Observations With Certificate
* [GetHosts](#gethosts) - Get multiple hosts
* [GetHost](#gethost) - Get a host
* [GetHostTimeline](#gethosttimeline) - Get host event history
* [GetWebProperties](#getwebproperties) - Get multiple web properties
* [GetWebProperty](#getwebproperty) - Get a web property
* [CreateTrackedScan](#createtrackedscan) - Create a tracked rescan
* [GetTrackedScan](#gettrackedscan) - Get tracked scan details
* [Aggregate](#aggregate) - Aggregate results for a search query
* [Search](#search) - Run a search query
* [GetTrackedScanThreatHunting](#gettrackedscanthreathunting) - Get tracked scan details

## GetCertificates

Retrieve information about multiple certificates. A certificate ID is its SHA-256 fingerprint in the Censys dataset.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-asset-certificate-list" method="get" path="/v3/global/asset/certificate" -->
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

    res, err := s.GlobalData.GetCertificates(ctx, operations.V3GlobaldataAssetCertificateListRequest{
        CertificateIds: []string{
            "3daf2843a77b6f4e6af43cd9b6f6746053b8c928e056e8a724808db8905a94cf",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeListCertificateAsset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.V3GlobaldataAssetCertificateListRequest](../../models/operations/v3globaldataassetcertificatelistrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.V3GlobaldataAssetCertificateListResponse](../../models/operations/v3globaldataassetcertificatelistresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## GetCertificate

Retrieve information about a single certificate. A certificate ID is its SHA-256 fingerprint in the Censys dataset.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-asset-certificate" method="get" path="/v3/global/asset/certificate/{certificate_id}" -->
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

    res, err := s.GlobalData.GetCertificate(ctx, operations.V3GlobaldataAssetCertificateRequest{
        CertificateID: "3daf2843a77b6f4e6af43cd9b6f6746053b8c928e056e8a724808db8905a94cf",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeCertificateAsset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.V3GlobaldataAssetCertificateRequest](../../models/operations/v3globaldataassetcertificaterequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.V3GlobaldataAssetCertificateResponse](../../models/operations/v3globaldataassetcertificateresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## GetHostObservationsWithCertificate

Retrieve historical observations of hosts associated with a certificate fingerprint. Useful for threat hunting, detection engineering, and timeline generation.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-get-host-observations-with-certificate" method="get" path="/v3/global/asset/certificate/{certificate_id}/observations/hosts" -->
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

    res, err := s.GlobalData.GetHostObservationsWithCertificate(ctx, operations.V3GlobaldataGetHostObservationsWithCertificateRequest{
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

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.V3GlobaldataGetHostObservationsWithCertificateRequest](../../models/operations/v3globaldatagethostobservationswithcertificaterequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.V3GlobaldataGetHostObservationsWithCertificateResponse](../../models/operations/v3globaldatagethostobservationswithcertificateresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## GetHosts

Retrieve information about multiple hosts. A host ID is its IP address.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-asset-host-list" method="get" path="/v3/global/asset/host" -->
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

    res, err := s.GlobalData.GetHosts(ctx, operations.V3GlobaldataAssetHostListRequest{
        HostIds: []string{
            "8.8.8.8",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeListHostAsset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.V3GlobaldataAssetHostListRequest](../../models/operations/v3globaldataassethostlistrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.V3GlobaldataAssetHostListResponse](../../models/operations/v3globaldataassethostlistresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## GetHost

Retrieve information about a single host. A host ID is its IP address.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-asset-host" method="get" path="/v3/global/asset/host/{host_id}" -->
```go
package main

import(
	"context"
	censyssdkgo "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/types"
	"github.com/censys/censys-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := censyssdkgo.New(
        censyssdkgo.WithOrganizationID("11111111-2222-3333-4444-555555555555"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.GlobalData.GetHost(ctx, operations.V3GlobaldataAssetHostRequest{
        HostID: "8.8.8.8",
        AtTime: types.MustNewTimeFromString("2025-01-01T00:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeHostAsset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.V3GlobaldataAssetHostRequest](../../models/operations/v3globaldataassethostrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.V3GlobaldataAssetHostResponse](../../models/operations/v3globaldataassethostresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## GetHostTimeline

Retrieve event history for a host. A host ID is its IP address.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-asset-host-timeline" method="get" path="/v3/global/asset/host/{host_id}/timeline" -->
```go
package main

import(
	"context"
	censyssdkgo "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/types"
	"github.com/censys/censys-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := censyssdkgo.New(
        censyssdkgo.WithOrganizationID("11111111-2222-3333-4444-555555555555"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.GlobalData.GetHostTimeline(ctx, operations.V3GlobaldataAssetHostTimelineRequest{
        HostID: "8.8.8.8",
        StartTime: types.MustTimeFromString("2025-01-01T00:00:00Z"),
        EndTime: types.MustTimeFromString("2025-01-02T00:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeHostTimeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.V3GlobaldataAssetHostTimelineRequest](../../models/operations/v3globaldataassethosttimelinerequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.V3GlobaldataAssetHostTimelineResponse](../../models/operations/v3globaldataassethosttimelineresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## GetWebProperties

Retrieve information about multiple web properties. Web properties are identified using a combination of a hostname and port joined with a colon, such as `platform.censys.io:80`.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-asset-webproperty-list" method="get" path="/v3/global/asset/webproperty" -->
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

    res, err := s.GlobalData.GetWebProperties(ctx, operations.V3GlobaldataAssetWebpropertyListRequest{
        WebpropertyIds: []string{
            "platform.censys.io:80",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeListWebpropertyAsset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.V3GlobaldataAssetWebpropertyListRequest](../../models/operations/v3globaldataassetwebpropertylistrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.V3GlobaldataAssetWebpropertyListResponse](../../models/operations/v3globaldataassetwebpropertylistresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## GetWebProperty

Retrieve information about a single web property. Web properties are identified using a combination of a hostname and port joined with a colon, such as `platform.censys.io:80`.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-asset-webproperty" method="get" path="/v3/global/asset/webproperty/{webproperty_id}" -->
```go
package main

import(
	"context"
	censyssdkgo "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/types"
	"github.com/censys/censys-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := censyssdkgo.New(
        censyssdkgo.WithOrganizationID("11111111-2222-3333-4444-555555555555"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.GlobalData.GetWebProperty(ctx, operations.V3GlobaldataAssetWebpropertyRequest{
        WebpropertyID: "platform.censys.io:80",
        AtTime: types.MustNewTimeFromString("2025-01-01T00:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeWebpropertyAsset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.V3GlobaldataAssetWebpropertyRequest](../../models/operations/v3globaldataassetwebpropertyrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.V3GlobaldataAssetWebpropertyResponse](../../models/operations/v3globaldataassetwebpropertyresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## CreateTrackedScan

Create a new tracked rescan for a known service or web property. Rescans are used to update information for previously discovered targets. The scan will be queued. The response will contain a scan ID that you can use with the [get tracked scan details endpoint](https://docs.censys.com/reference/v3-globaldata-scans-get#/) to monitor its status and results.<br><br>This endpoint is available to all Enterprise customers.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-scans-rescan" method="post" path="/v3/global/scans/rescan" -->
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

    res, err := s.GlobalData.CreateTrackedScan(ctx, operations.V3GlobaldataScansRescanRequest{
        ScansRescanInputBody: components.ScansRescanInputBody{
            Target: components.CreateScansRescanInputBodyTargetTwo(
                components.Two{
                    WebOrigin: components.TargetWebOrigin{
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.V3GlobaldataScansRescanRequest](../../models/operations/v3globaldatascansrescanrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.V3GlobaldataScansRescanResponse](../../models/operations/v3globaldatascansrescanresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

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

    res, err := s.GlobalData.GetTrackedScan(ctx, operations.V3GlobaldataScansGetRequest{
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

## Aggregate

Aggregate results for a Platform search query. This functionality is equivalent to the [Report Builder](https://docs.censys.com/docs/platform-report-builder#/) in the Platform web UI.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-search-aggregate" method="post" path="/v3/global/search/aggregate" -->
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

    res, err := s.GlobalData.Aggregate(ctx, operations.V3GlobaldataSearchAggregateRequest{
        SearchAggregateInputBody: components.SearchAggregateInputBody{
            Field: "host.services.port",
            NumberOfBuckets: 100,
            Query: "host.services.protocol=SSH",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeSearchAggregateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.V3GlobaldataSearchAggregateRequest](../../models/operations/v3globaldatasearchaggregaterequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.V3GlobaldataSearchAggregateResponse](../../models/operations/v3globaldatasearchaggregateresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## Search

Run a search query across Censys data. Reference the [documentation on Censys Query Language](https://docs.censys.com/docs/censys-query-language#/) for information about query syntax.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-search-query" method="post" path="/v3/global/search/query" -->
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

    res, err := s.GlobalData.Search(ctx, operations.V3GlobaldataSearchQueryRequest{
        SearchQueryInputBody: components.SearchQueryInputBody{
            Fields: []string{
                "host.ip",
            },
            PageSize: censyssdkgo.Int64(1),
            Query: "host.services: (protocol=SSH and not port: 22)",
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

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.V3GlobaldataSearchQueryRequest](../../models/operations/v3globaldatasearchqueryrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.V3GlobaldataSearchQueryResponse](../../models/operations/v3globaldatasearchqueryresponse.md), error**

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

    res, err := s.GlobalData.GetTrackedScanThreatHunting(ctx, operations.V3ThreathuntingScansGetRequest{
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