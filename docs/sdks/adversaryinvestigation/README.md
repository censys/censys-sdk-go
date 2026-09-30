# AdversaryInvestigation

## Overview

Endpoints related to the Adversary Investigation product

### Available Operations

* [CreateCenseyeJob](#createcenseyejob) - CensEye: Create a pivot analysis job
* [GetCenseyeJob](#getcenseyejob) - CensEye: Get job status
* [GetCenseyeJobResults](#getcenseyejobresults) - CensEye: Get job results
* [GetHostObservationsWithCertificate](#gethostobservationswithcertificate) - Get host history for a certificate
* [CreateInvestigationFileUpload](#createinvestigationfileupload) - Investigations: Create file upload
* [ListInvestigationJobs](#listinvestigationjobs) - Investigations: List jobs
* [CreateInvestigationJob](#createinvestigationjob) - Investigations: Create job
* [GetInvestigationJob](#getinvestigationjob) - Investigations: Get job status
* [GetInvestigationJobResults](#getinvestigationjobresults) - Investigations: Get job results
* [GetInvestigationUsage](#getinvestigationusage) - Investigations: Get usage
* [CreateTrackedScan](#createtrackedscan) - Live Discovery: Initiate a new scan
* [ListThreats](#listthreats) - List active threats
* [ValueCounts](#valuecounts) - CensEye: Retrieve value counts to discover pivots

## CreateCenseyeJob

Create an asynchronous CensEye pivot analysis job for a host, web property, or certificate. The job extracts [default pivot fields](https://docs.censys.com/docs/platform-threat-hunting-use-censeye-to-build-detections#default-pivot-fields) from the target asset and counts matching documents for each field-value pair. Poll the job status endpoint to track progress, then retrieve results when complete.<br><br>To use this endpoint, your organization must have access to the Adversary Investigation module.<br><br>This endpoint costs 44 credits to execute for a host, 28 credits to execute for a web property, and 7 credits to execute for a certificate.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-threathunting-censeye-jobs-create" method="post" path="/v3/threat-hunting/censeye/jobs" -->
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

    res, err := s.AdversaryInvestigation.CreateCenseyeJob(ctx, operations.V3ThreathuntingCenseyeJobsCreateRequest{
        CreateCenseyeJobInputBody: components.CreateCenseyeJobInputBody{
            Target: components.CenseyeTarget{
                CertificateID: censyssdkgo.Pointer("3daf2843a77b6f4e6af43cd9b6f6746053b8c928e056e8a724808db8905a94cf"),
                HostID: censyssdkgo.Pointer("8.8.8.8"),
                WebpropertyID: censyssdkgo.Pointer("example.com:443"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeCenseyeJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.V3ThreathuntingCenseyeJobsCreateRequest](../../models/operations/v3threathuntingcenseyejobscreaterequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.V3ThreathuntingCenseyeJobsCreateResponse](../../models/operations/v3threathuntingcenseyejobscreateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 400, 403, 422                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetCenseyeJob

Retrieve the current status of a CensEye pivot analysis job. Use this to poll for completion before fetching results.<br><br>To use this endpoint, your organization must have access to the Adversary Investigation module.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-threathunting-censeye-jobs-get" method="get" path="/v3/threat-hunting/censeye/jobs/{job_id}" -->
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
        censyssdkgo.WithOrganizationID("<id>"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AdversaryInvestigation.GetCenseyeJob(ctx, operations.V3ThreathuntingCenseyeJobsGetRequest{
        JobID: "3c47b971-5db6-4a9e-8d59-14fc0486172b",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeCenseyeJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.V3ThreathuntingCenseyeJobsGetRequest](../../models/operations/v3threathuntingcenseyejobsgetrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.V3ThreathuntingCenseyeJobsGetResponse](../../models/operations/v3threathuntingcenseyejobsgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 400, 403, 404                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetCenseyeJobResults

Retrieve the results of a completed CensEye pivot analysis job. Each result contains a count and the field-value pairs that were analyzed. Results may be empty if the job is still running.<br><br>Results are paginated. Use the `next_page_token` from the response to fetch subsequent pages.<br><br>To use this endpoint, your organization must have access to the Adversary Investigation module.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-threathunting-censeye-job-results" method="get" path="/v3/threat-hunting/censeye/jobs/{job_id}/results" -->
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
        censyssdkgo.WithOrganizationID("<id>"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AdversaryInvestigation.GetCenseyeJobResults(ctx, operations.V3ThreathuntingCenseyeJobResultsRequest{
        JobID: "e58e9a0e-e104-42cf-9d0e-fe88713bc6e3",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeCenseyeResultsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.V3ThreathuntingCenseyeJobResultsRequest](../../models/operations/v3threathuntingcenseyejobresultsrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.V3ThreathuntingCenseyeJobResultsResponse](../../models/operations/v3threathuntingcenseyejobresultsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 400, 403, 404                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetHostObservationsWithCertificate

Retrieve the historical observations of hosts associated with a certificate. This is useful for threat hunting, detection engineering, and timeline generation. Certificate history is also visible to Adversary Investigation users in the Platform UI on the [certificate timeline](https://docs.censys.com/docs/platform-threat-hunting-use-cert-history-to-build-better-detections#/).<br><br>You can define a specific time frame of interest. If you do not specify a time frame, this endpoint will search the historical dataset that is available to your account.<br><br>For workspaces with a history limit, the API returns observation ranges that overlap the allowed history window. The window starts at 00:00 UTC the allowed number of days ago and ends at the time of the request. Ranges are returned in full, with their original start and end times, even if they extend outside the window.<br><br>You may also filter results by port and transport protocol.<br><br>This endpoint is available to organizations that have access to the Adversary Investigation module. This endpoint costs one credit per page of results.

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
        censyssdkgo.WithOrganizationID("<id>"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AdversaryInvestigation.GetHostObservationsWithCertificate(ctx, operations.V3ThreathuntingGetHostObservationsWithCertificateRequest{
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
| sdkerrors.ErrorModel          | 400, 403, 404                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## CreateInvestigationFileUpload

Prepare to attach an evidence file to an investigation. This endpoint does not accept the file. It returns a URL to send the file to and an identifier to pass to the [create job endpoint](https://docs.censys.com/reference/v3-threathunting-investigations-jobs-create) afterwards.<br><br>Upload the file with an HTTP PUT to `upload_url` and send `upload_headers` exactly as returned by this endpoint. The URL grants access to that one file and stops working at `expire_time`; a file uploaded before that stays usable afterwards. Prepare one upload per file.<br><br>Start the investigation promptly after uploading. There is a limit on how many uploaded files you may hold without using them and starting an investigation from a file releases its slot. A file that no investigation references is discarded once it passes the service's retention window for unused uploads. A file an investigation does reference is kept with that investigation for as long as the investigation itself.<br><br>To use this endpoint, your organization must have access to the Adversary Investigation module.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-threathunting-investigations-files-create" method="post" path="/v3/threat-hunting/investigations/files" -->
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

    res, err := s.AdversaryInvestigation.CreateInvestigationFileUpload(ctx, operations.V3ThreathuntingInvestigationsFilesCreateRequest{
        CreateInvestigationFileInputBody: components.CreateInvestigationFileInputBody{
            ContentType: components.ContentTypeApplicationPdf,
            Filename: censyssdkgo.Pointer("incident-report.pdf"),
            SizeBytes: 1048576,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeInvestigationFileUpload != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.V3ThreathuntingInvestigationsFilesCreateRequest](../../models/operations/v3threathuntinginvestigationsfilescreaterequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.V3ThreathuntingInvestigationsFilesCreateResponse](../../models/operations/v3threathuntinginvestigationsfilescreateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 409, 413, 422, 429       | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## ListInvestigationJobs

List the AI investigations you have started. The most recent investigations are listed first and only investigations that have not passed their retention period are shown. Results are paginated and include investigations started in the Platform UI as well as through the API.<br><br>To use this endpoint, your organization must have access to the Adversary Investigation module.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-threathunting-investigations-jobs-list" method="get" path="/v3/threat-hunting/investigations/jobs" -->
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
        censyssdkgo.WithOrganizationID("<id>"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AdversaryInvestigation.ListInvestigationJobs(ctx, operations.V3ThreathuntingInvestigationsJobsListRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeInvestigationJobsList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.V3ThreathuntingInvestigationsJobsListRequest](../../models/operations/v3threathuntinginvestigationsjobslistrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.V3ThreathuntingInvestigationsJobsListResponse](../../models/operations/v3threathuntinginvestigationsjobslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 409, 422                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## CreateInvestigationJob

Start an [AI investigation](https://docs.censys.com/docs/platform-ai-investigations) from a set of indicators, a set of previously uploaded evidence files, or both. Supply at least one `indicators` or `file_ids`. To use files, upload them first with the [create file upload endpoint](https://docs.censys.com/reference/v3-threathunting-investigations-files-create).<br><br>This endpoint returns a `job_id` that you can poll to retrieve its status and results.<br><br>Provide `start_time` and `end_time` to scope the investigation to a time frame, or `start_time` alone to scope it from that time up to now. Omit both to investigate current data.<br><br>To use this endpoint, your organization must have access to the Adversary Investigation module.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-threathunting-investigations-jobs-create" method="post" path="/v3/threat-hunting/investigations/jobs" -->
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
        censyssdkgo.WithOrganizationID("<id>"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AdversaryInvestigation.CreateInvestigationJob(ctx, operations.V3ThreathuntingInvestigationsJobsCreateRequest{
        CreateInvestigationJobInputBody: components.CreateInvestigationJobInputBody{
            EndTime: types.MustNewTimeFromString("2026-08-01T00:00:00Z"),
            FileIds: []string{
                "file_550e8400-e29b-41d4-a716-446655440000",
            },
            Indicators: []string{
                "1.1.1.1",
                "example.com",
            },
            StartTime: types.MustNewTimeFromString("2026-06-01T00:00:00Z"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeCreatedInvestigationJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.V3ThreathuntingInvestigationsJobsCreateRequest](../../models/operations/v3threathuntinginvestigationsjobscreaterequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.V3ThreathuntingInvestigationsJobsCreateResponse](../../models/operations/v3threathuntinginvestigationsjobscreateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 409, 413, 422, 429       | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetInvestigationJob

Retrieve the status of one AI investigation. Poll this endpoint until the investigation is completed, then download its report and evidence.<br><br>An investigation that does not exist, belongs to another user, or has passed its retention window will return a “not found” response.<br><br>To use this endpoint, your organization must have access to the Adversary Investigation module.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-threathunting-investigations-jobs-get" method="get" path="/v3/threat-hunting/investigations/jobs/{job_id}" -->
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
        censyssdkgo.WithOrganizationID("<id>"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AdversaryInvestigation.GetInvestigationJob(ctx, operations.V3ThreathuntingInvestigationsJobsGetRequest{
        JobID: "9f1c7f36-34af-416b-82b3-b496756f4d5e",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeInvestigationJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.V3ThreathuntingInvestigationsJobsGetRequest](../../models/operations/v3threathuntinginvestigationsjobsgetrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.V3ThreathuntingInvestigationsJobsGetResponse](../../models/operations/v3threathuntinginvestigationsjobsgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 409, 422            | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetInvestigationJobResults

Download the ZIP archive that contains a completed AI investigation's report and evidence. You can only retrieve the job results for an investigation you started. The archive is composed when you request it and is never stored.<br><br>Investigations that did not publish any findings or have passed their retention windows will return a “not found” response.<br><br>If an investigation’s ZIP archive is larger than 20 megabytes, you will receive a “payload too large” response. You can only retrieve the results for investigations that exceed this size limit within the Platform UI.<br><br>To use this endpoint, your organization must have access to the Adversary Investigation module.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-threathunting-investigations-jobs-results" method="get" path="/v3/threat-hunting/investigations/jobs/{job_id}/results" -->
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
        censyssdkgo.WithOrganizationID("<id>"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AdversaryInvestigation.GetInvestigationJobResults(ctx, operations.V3ThreathuntingInvestigationsJobsResultsRequest{
        JobID: "abfb0cd9-3719-491a-be31-8dbbaaff559c",
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

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.V3ThreathuntingInvestigationsJobsResultsRequest](../../models/operations/v3threathuntinginvestigationsjobsresultsrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.V3ThreathuntingInvestigationsJobsResultsResponse](../../models/operations/v3threathuntinginvestigationsjobsresultsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 409, 413, 422       | application/problem+json      |
| sdkerrors.ErrorModel          | 500, 503, 504                 | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetInvestigationUsage

Retrieve your organization's investigation limit, current usage, and the number of remaining investigations.<br><br>To use this endpoint, your organization must have access to the Adversary Investigation module.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-threathunting-investigations-usage-get" method="get" path="/v3/threat-hunting/investigations/usage" -->
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
        censyssdkgo.WithOrganizationID("<id>"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AdversaryInvestigation.GetInvestigationUsage(ctx, operations.V3ThreathuntingInvestigationsUsageGetRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeInvestigationUsage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.V3ThreathuntingInvestigationsUsageGetRequest](../../models/operations/v3threathuntinginvestigationsusagegetrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.V3ThreathuntingInvestigationsUsageGetResponse](../../models/operations/v3threathuntinginvestigationsusagegetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 409, 422                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## CreateTrackedScan

Initiate a scan to look for a currently unobserved service at a specific IP and port (`ip:port`) or hostname and port (`hostname:port`). This is equivalent to the [Live Discovery](https://docs.censys.com/docs/platform-threat-hunting-use-live-scan-and-rescan-to-validate-infrastructure#/) feature available in the UI, but you can also target web properties in addition to hosts.<br><br>The scan may take several minutes to complete. The response will contain a scan ID that you can use to [monitor the scan's status](https://docs.censys.com/reference/v3-threathunting-scans-get#/). After the scan completes, perform a lookup on the target asset to retrieve detailed scan information.<br><br>This endpoint is available to organizations that have access to the Adversary Investigation module. It costs 15 credits to execute this endpoint.

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
        censyssdkgo.WithOrganizationID("<id>"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AdversaryInvestigation.CreateTrackedScan(ctx, operations.V3ThreathuntingScansDiscoveryRequest{
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
| sdkerrors.ErrorModel          | 400, 403, 422                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## ListThreats

Retrieve a list of active threats observed by Censys by aggregating threat IDs across hosts and web properties. Threats are active if their fingerprint has been identified on hosts or web properties by Censys scans. This information is also available on the [Explore Threats page in the Platform web UI](https://platform.censys.io/threats).<br><br>This endpoint is available to organizations that have access to the Adversary Investigation module.

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
        censyssdkgo.WithOrganizationID("<id>"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AdversaryInvestigation.ListThreats(ctx, operations.V3ThreathuntingThreatsListRequest{
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
| sdkerrors.ErrorModel          | 400, 403, 422                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## ValueCounts

Get counts of web assets for specific field-value pairs and combinations of field-value pairs. This is similar to the [CensEye functionality](https://docs.censys.com/docs/platform-threat-hunting-use-censeye-to-build-detections#/) available in the Platform web UI, but it allows you to define specific fields of interest rather than the [default fields](https://docs.censys.com/docs/platform-threat-hunting-use-censeye-to-build-detections#default-pivot-fields) leveraged by the tool in the UI.<br><br>Each array can only target fields within the same nested object and may contain at most 5 field-value pairs. For example, you can combine `host.services.port=80` and `host.services.protocol=SSH` in the same array, but you cannot combine `host.services.port=80` and `host.location.country="United States"` in the same array. You can input multiple arrays of objects in each API call.<br><br>To use this endpoint, your organization must have access to the Adversary Investigation module. This endpoint costs 1 credit per count condition (array of objects) included in the API call.

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
        censyssdkgo.WithOrganizationID("<id>"),
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AdversaryInvestigation.ValueCounts(ctx, operations.V3ThreathuntingValueCountsRequest{
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
| sdkerrors.ErrorModel          | 400, 403, 422                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |