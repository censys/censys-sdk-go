# github.com/censys/censys-sdk-go

Developer-friendly & type-safe Go SDK specifically catered to leverage *openapi* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=openapi&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<!-- Start Summary [summary] -->
## Summary


<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/censys/censys-sdk-go](#githubcomcensyscensys-sdk-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Global Parameters](#global-parameters)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Authentication](#authentication)
  * [Special Types](#special-types)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/censys/censys-sdk-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AccountManagement](docs/sdks/accountmanagement/README.md)

* [GetOrganizationDetails](docs/sdks/accountmanagement/README.md#getorganizationdetails) - Get organization details
* [GetOrganizationCredits](docs/sdks/accountmanagement/README.md#getorganizationcredits) - Get organization credit balance
* [GetOrganizationCreditUsage](docs/sdks/accountmanagement/README.md#getorganizationcreditusage) - Get organization credit usage
* [InviteUserToOrganization](docs/sdks/accountmanagement/README.md#inviteusertoorganization) - Invite user to organization
* [ListOrganizationMembers](docs/sdks/accountmanagement/README.md#listorganizationmembers) - List organization members
* [RemoveOrganizationMember](docs/sdks/accountmanagement/README.md#removeorganizationmember) - Remove member from organization
* [UpdateOrganizationMember](docs/sdks/accountmanagement/README.md#updateorganizationmember) - Update a member's roles in an organization
* [GetMemberCreditUsage](docs/sdks/accountmanagement/README.md#getmembercreditusage) - Get organization member credit usage
* [GetUserCredits](docs/sdks/accountmanagement/README.md#getusercredits) - Get Free user credit balance
* [GetUserCreditsUsage](docs/sdks/accountmanagement/README.md#getusercreditsusage) - Get Free user credit usage

### [Collections](docs/sdks/collections/README.md)

* [List](docs/sdks/collections/README.md#list) - List collections
* [Create](docs/sdks/collections/README.md#create) - Create a collection
* [Delete](docs/sdks/collections/README.md#delete) - Delete a collection
* [Get](docs/sdks/collections/README.md#get) - Get a collection
* [Update](docs/sdks/collections/README.md#update) - Update a collection
* [ListEvents](docs/sdks/collections/README.md#listevents) - Get a collection's events
* [Aggregate](docs/sdks/collections/README.md#aggregate) - Aggregate results for a search query within a collection
* [Search](docs/sdks/collections/README.md#search) - Run a search query within a collection

### [GlobalData](docs/sdks/globaldata/README.md)

* [GetCertificates](docs/sdks/globaldata/README.md#getcertificates) - Retrieve multiple certificates
* [GetCertificatesRaw](docs/sdks/globaldata/README.md#getcertificatesraw) - Retrieve multiple certificates in PEM format
* [GetCertificate](docs/sdks/globaldata/README.md#getcertificate) - Get a certificate
* [GetCertificateRaw](docs/sdks/globaldata/README.md#getcertificateraw) - Get a certificate in PEM format
* [GetHosts](docs/sdks/globaldata/README.md#gethosts) - Retrieve multiple hosts
* [GetHost](docs/sdks/globaldata/README.md#gethost) - Get a host
* [GetHostTimeline](docs/sdks/globaldata/README.md#gethosttimeline) - Get host event history
* [GetWebProperties](docs/sdks/globaldata/README.md#getwebproperties) - Retrieve multiple web properties
* [GetWebProperty](docs/sdks/globaldata/README.md#getwebproperty) - Get a web property
* [CreateTrackedScan](docs/sdks/globaldata/README.md#createtrackedscan) - Live Rescan: Initiate a new rescan
* [GetTrackedScan](docs/sdks/globaldata/README.md#gettrackedscan) - Get scan status
* [Aggregate](docs/sdks/globaldata/README.md#aggregate) - Aggregate results for a search query
* [ConvertLegacySearchQueries](docs/sdks/globaldata/README.md#convertlegacysearchqueries) - Convert Legacy Search queries to Platform queries
* [Search](docs/sdks/globaldata/README.md#search) - Run a search query

### [ThreatHunting](docs/sdks/threathunting/README.md)

* [GetHostObservationsWithCertificate](docs/sdks/threathunting/README.md#gethostobservationswithcertificate) - Get host history for a certificate
* [CreateTrackedScan](docs/sdks/threathunting/README.md#createtrackedscan) - Live Discovery: Initiate a new scan
* [GetTrackedScanThreatHunting](docs/sdks/threathunting/README.md#gettrackedscanthreathunting) - Get scan status
* [ListThreats](docs/sdks/threathunting/README.md#listthreats) - List active threats
* [ValueCounts](docs/sdks/threathunting/README.md#valuecounts) - CensEye: Retrieve value counts to discover pivots

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Global Parameters [global-parameters] -->
## Global Parameters

A parameter is configured globally. This parameter may be set on the SDK client instance itself during initialization. When configured as an option during SDK initialization, This global value will be used as the default on the operations that use it. When such operations are called, there is a place in each to override the global value, if needed.

For example, you can set `organization_id` to `"<id>"` at SDK initialization and then you do not have to pass the same value on calls to operations like `GetOrganizationDetails`. But if you want to do so you may, which will locally override the global setting. See the example code below for a demonstration.


### Available Globals

The following global parameter is available.

| Name           | Type   | Description                   |
| -------------- | ------ | ----------------------------- |
| OrganizationID | string | The OrganizationID parameter. |

### Example

```go
package main

import (
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

	res, err := s.AccountManagement.GetOrganizationDetails(ctx, operations.V3AccountmanagementOrgDetailsRequest{
		OrganizationID: "11111111-2222-3333-4444-555555555555",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseEnvelopeOrganizationDetails != nil {
		// handle response
	}
}

```
<!-- End Global Parameters [global-parameters] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	censyssdkgo "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/models/operations"
	"github.com/censys/censys-sdk-go/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := censyssdkgo.New(
		censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.AccountManagement.GetOrganizationDetails(ctx, operations.V3AccountmanagementOrgDetailsRequest{
		OrganizationID: "11111111-2222-3333-4444-555555555555",
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseEnvelopeOrganizationDetails != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	censyssdkgo "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/models/operations"
	"github.com/censys/censys-sdk-go/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := censyssdkgo.New(
		censyssdkgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.AccountManagement.GetOrganizationDetails(ctx, operations.V3AccountmanagementOrgDetailsRequest{
		OrganizationID: "11111111-2222-3333-4444-555555555555",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseEnvelopeOrganizationDetails != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetOrganizationDetails` function may return the following errors:

| Error Type                    | Status Code   | Content Type             |
| ----------------------------- | ------------- | ------------------------ |
| sdkerrors.AuthenticationError | 401           | application/json         |
| sdkerrors.ErrorModel          | 403, 404, 422 | application/problem+json |
| sdkerrors.SDKError            | 4XX, 5XX      | \*/\*                    |

### Example

```go
package main

import (
	"context"
	"errors"
	censyssdkgo "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/models/operations"
	"github.com/censys/censys-sdk-go/models/sdkerrors"
	"log"
)

func main() {
	ctx := context.Background()

	s := censyssdkgo.New(
		censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.AccountManagement.GetOrganizationDetails(ctx, operations.V3AccountmanagementOrgDetailsRequest{
		OrganizationID: "11111111-2222-3333-4444-555555555555",
	})
	if err != nil {

		var e *sdkerrors.AuthenticationError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ErrorModel
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	censyssdkgo "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := censyssdkgo.New(
		censyssdkgo.WithServerURL("https://api.platform.censys.io"),
		censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.AccountManagement.GetOrganizationDetails(ctx, operations.V3AccountmanagementOrgDetailsRequest{
		OrganizationID: "11111111-2222-3333-4444-555555555555",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseEnvelopeOrganizationDetails != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/censys/censys-sdk-go"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = censyssdkgo.New(censyssdkgo.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name                  | Type | Scheme      |
| --------------------- | ---- | ----------- |
| `PersonalAccessToken` | http | HTTP Bearer |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	censyssdkgo "github.com/censys/censys-sdk-go"
	"github.com/censys/censys-sdk-go/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := censyssdkgo.New(
		censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.AccountManagement.GetOrganizationDetails(ctx, operations.V3AccountmanagementOrgDetailsRequest{
		OrganizationID: "11111111-2222-3333-4444-555555555555",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseEnvelopeOrganizationDetails != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation.
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release.

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=openapi&utm_campaign=go)
