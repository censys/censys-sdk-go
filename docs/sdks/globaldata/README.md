# GlobalData
(*GlobalData*)

## Overview

Endpoints related to the Global Data product

### Available Operations

* [GetCertificates](#getcertificates) - Retrieve multiple certificates
* [GetCertificatesRaw](#getcertificatesraw) - Retrieve multiple certificates in PEM format
* [GetCertificate](#getcertificate) - Get a certificate
* [GetCertificateRaw](#getcertificateraw) - Get a certificate in PEM format
* [GetHosts](#gethosts) - Retrieve multiple hosts
* [GetHost](#gethost) - Get a host
* [GetHostTimeline](#gethosttimeline) - Get host event history
* [GetWebProperties](#getwebproperties) - Retrieve multiple web properties
* [GetWebProperty](#getwebproperty) - Get a web property
* [CreateTrackedScan](#createtrackedscan) - Live Rescan: Initiate a new rescan
* [GetTrackedScan](#gettrackedscan) - Get scan status
* [Aggregate](#aggregate) - Aggregate results for a search query
* [ConvertLegacySearchQueries](#convertlegacysearchqueries) - Convert Legacy Search queries to Platform queries
* [Search](#search) - Run a search query

## GetCertificates

Retrieve information about multiple certificates. A certificate ID is its SHA-256 fingerprint in the Censys dataset.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-asset-certificate-list-post" method="post" path="/v3/global/asset/certificate" -->
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

    res, err := s.GlobalData.GetCertificates(ctx, operations.V3GlobaldataAssetCertificateListPostRequest{
        AssetCertificateListInputBody: components.AssetCertificateListInputBody{
            CertificateIds: []string{
                "3daf2843a77b6f4e6af43cd9b6f6746053b8c928e056e8a724808db8905a94cf",
            },
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.V3GlobaldataAssetCertificateListPostRequest](../../models/operations/v3globaldataassetcertificatelistpostrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.V3GlobaldataAssetCertificateListPostResponse](../../models/operations/v3globaldataassetcertificatelistpostresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404                      | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetCertificatesRaw

Retrieve the raw PEM-encoded format for multiple certificates. A certificate ID is its SHA-256 fingerprint in the Censys dataset.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-asset-certificate-list-raw-post" method="post" path="/v3/global/asset/certificate/raw" -->
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

    res, err := s.GlobalData.GetCertificatesRaw(ctx, operations.V3GlobaldataAssetCertificateListRawPostRequest{
        AssetCertificateListInputBody: components.AssetCertificateListInputBody{
            CertificateIds: []string{
                "3daf2843a77b6f4e6af43cd9b6f6746053b8c928e056e8a724808db8905a94cf",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeListRawCertificateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.V3GlobaldataAssetCertificateListRawPostRequest](../../models/operations/v3globaldataassetcertificatelistrawpostrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.V3GlobaldataAssetCertificateListRawPostResponse](../../models/operations/v3globaldataassetcertificatelistrawpostresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404                      | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

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

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404                      | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetCertificateRaw

Retrieve the raw PEM-encoded format of a certificate. A certificate ID is its SHA-256 fingerprint in the Censys dataset.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-asset-certificate-raw" method="get" path="/v3/global/asset/certificate/{certificate_id}/raw" -->
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

    res, err := s.GlobalData.GetCertificateRaw(ctx, operations.V3GlobaldataAssetCertificateRawRequest{
        CertificateID: "3daf2843a77b6f4e6af43cd9b6f6746053b8c928e056e8a724808db8905a94cf",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.V3GlobaldataAssetCertificateRawRequest](../../models/operations/v3globaldataassetcertificaterawrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.V3GlobaldataAssetCertificateRawResponse](../../models/operations/v3globaldataassetcertificaterawresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404                      | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetHosts

Retrieve information about multiple hosts. A host ID is its IP address.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-asset-host-list-post" method="post" path="/v3/global/asset/host" -->
```go
package main

import(
	"context"
	censyssdkgo "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/types"
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

    res, err := s.GlobalData.GetHosts(ctx, operations.V3GlobaldataAssetHostListPostRequest{
        AssetHostListInputBody: components.AssetHostListInputBody{
            AtTime: types.MustNewTimeFromString("2025-01-01T00:00:00Z"),
            HostIds: []string{
                "8.8.8.8",
            },
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.V3GlobaldataAssetHostListPostRequest](../../models/operations/v3globaldataassethostlistpostrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.V3GlobaldataAssetHostListPostResponse](../../models/operations/v3globaldataassethostlistpostresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

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

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404                      | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetHostTimeline

Retrieve event history for a host. A host ID is its IP address.<br><br>Note that when a service protocol changes after a new scan (for example, from `UNKNOWN` to `NETBIOS`), this information will only be reflected in the `scan` object. It will not be shown in the `service_scanned diff` object.

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
        StartTime: types.MustTimeFromString("2025-01-02T00:00:00Z"),
        EndTime: types.MustTimeFromString("2025-01-01T00:00:00Z"),
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

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetWebProperties

Retrieve information about multiple web properties. Web properties are identified using a combination of a hostname and port joined with a colon, such as `platform.censys.io:80`.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-asset-webproperty-list-post" method="post" path="/v3/global/asset/webproperty" -->
```go
package main

import(
	"context"
	censyssdkgo "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/types"
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

    res, err := s.GlobalData.GetWebProperties(ctx, operations.V3GlobaldataAssetWebpropertyListPostRequest{
        AssetWebpropertyListInputBody: components.AssetWebpropertyListInputBody{
            AtTime: types.MustNewTimeFromString("2025-01-01T00:00:00Z"),
            WebpropertyIds: []string{
                "platform.censys.io:80",
            },
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.V3GlobaldataAssetWebpropertyListPostRequest](../../models/operations/v3globaldataassetwebpropertylistpostrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.V3GlobaldataAssetWebpropertyListPostResponse](../../models/operations/v3globaldataassetwebpropertylistpostresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

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

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 422                      | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## CreateTrackedScan

Initiate a rescan for a known host service at a specific IP and port (`ip:port`) or hostname and port (`hostname:port`). This is equivalent to the [Live Rescan](https://docs.censys.com/docs/platform-live-rescan#/) feature available in the UI, but you can also target web properties in addition to hosts.<br><br>The scan may take several minutes to complete. The response will contain a scan ID that you can use to [monitor the scan's status](https://docs.censys.com/reference/v3-globaldata-scans-get#/). After the scan completes, perform a lookup on the target asset to retrieve detailed scan information.<br><br>This endpoint is available to all Enterprise customers. It costs 10 credits to execute.

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

Retrieve the current status of a scan by its ID. This endpoint works for both [Live Discovery scans](https://docs.censys.com/reference/v3-threathunting-scans-discovery#/) and [Live Rescans](https://docs.censys.com/reference/v3-globaldata-scans-rescan#/).<br><br>If the scan was successful, perform a lookup on the target asset to retrieve detailed scan information.<br><br>This endpoint is available to all Enterprise customers. This endpoint does not cost any credits to execute.

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

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 422                      | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## ConvertLegacySearchQueries

Convert Censys Search Language queries used in Legacy Search into Censys Query Language (CenQL) queries for use in the Platform.<br><br>Reference the [documentation on CenQL](https://docs.censys.com/docs/censys-query-language) for more information about query syntax.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-globaldata-search-convert" method="post" path="/v3/global/search/convert" -->
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

    res, err := s.GlobalData.ConvertLegacySearchQueries(ctx, operations.V3GlobaldataSearchConvertRequest{
        SearchConvertQueryInputBody: components.SearchConvertQueryInputBody{
            Queries: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeListSearchConvertQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.V3GlobaldataSearchConvertRequest](../../models/operations/v3globaldatasearchconvertrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.V3GlobaldataSearchConvertResponse](../../models/operations/v3globaldatasearchconvertresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403                           | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## Search

Run a search query across Censys data. Reference the [documentation on Censys Query Language](https://docs.censys.com/docs/censys-query-language#/) for information about query syntax. Host services that match your search criteria will be returned in a `matched_services` object.

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
            PageSize: censyssdkgo.Pointer[int64](1),
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

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |