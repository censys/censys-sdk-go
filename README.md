# openapi

Developer-friendly & type-safe Go SDK specifically catered to leverage *openapi* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=openapi&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/censys/censys). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary


<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [openapi](#openapi)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Global Parameters](#global-parameters)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Authentication](#authentication)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/censys/censys-sdk-go-internal
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Collections](docs/sdks/collections/README.md)

* [List](docs/sdks/collections/README.md#list) - List Collections
* [Create](docs/sdks/collections/README.md#create) - Create a Collection
* [Delete](docs/sdks/collections/README.md#delete) - Delete a Collection
* [Get](docs/sdks/collections/README.md#get) - Get a Collection
* [Update](docs/sdks/collections/README.md#update) - Update a Collection
* [ListEvents](docs/sdks/collections/README.md#listevents) - List a Collection's events
* [Aggregate](docs/sdks/collections/README.md#aggregate) - Search / Aggregate
* [Search](docs/sdks/collections/README.md#search) - Search / Query

### [GlobalData](docs/sdks/globaldata/README.md)

* [GetCertificates](docs/sdks/globaldata/README.md#getcertificates) - Asset / Certificate Bulk
* [GetCertificate](docs/sdks/globaldata/README.md#getcertificate) - Asset / Certificate
* [GetHosts](docs/sdks/globaldata/README.md#gethosts) - Asset / Host Bulk
* [GetHost](docs/sdks/globaldata/README.md#gethost) - Asset / Host
* [GetHostTimeline](docs/sdks/globaldata/README.md#gethosttimeline) - Asset / Host Timeline
* [GetWebProperties](docs/sdks/globaldata/README.md#getwebproperties) - Asset / WebProperty Bulk
* [GetWebProperty](docs/sdks/globaldata/README.md#getwebproperty) - Asset / WebProperty
* [Aggregate](docs/sdks/globaldata/README.md#aggregate) - Search / Aggregate
* [Search](docs/sdks/globaldata/README.md#search) - Search / Query


</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Global Parameters [global-parameters] -->
## Global Parameters

A parameter is configured globally. This parameter may be set on the SDK client instance itself during initialization. When configured as an option during SDK initialization, This global value will be used as the default on the operations that use it. When such operations are called, there is a place in each to override the global value, if needed.

For example, you can set `organization_id` to `"<id>"` at SDK initialization and then you do not have to pass the same value on calls to operations like `List`. But if you want to do so you may, which will locally override the global setting. See the example code below for a demonstration.


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
	censyssdkgointernal "github.com/censys/censys-sdk-go-internal"
	"github.com/censys/censys-sdk-go-internal/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := censyssdkgointernal.New(
		censyssdkgointernal.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.Collections.List(ctx, operations.V3CollectionsCrudListRequest{
		OrganizationID: censyssdkgointernal.String("<id>"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseEnvelopeListCollectionsResponseV1 != nil {
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
	censyssdkgointernal "github.com/censys/censys-sdk-go-internal"
	"github.com/censys/censys-sdk-go-internal/models/operations"
	"github.com/censys/censys-sdk-go-internal/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := censyssdkgointernal.New(
		censyssdkgointernal.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.Collections.List(ctx, operations.V3CollectionsCrudListRequest{
		OrganizationID: censyssdkgointernal.String("<id>"),
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
	if res.ResponseEnvelopeListCollectionsResponseV1 != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	censyssdkgointernal "github.com/censys/censys-sdk-go-internal"
	"github.com/censys/censys-sdk-go-internal/models/operations"
	"github.com/censys/censys-sdk-go-internal/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := censyssdkgointernal.New(
		censyssdkgointernal.WithRetryConfig(
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
		censyssdkgointernal.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.Collections.List(ctx, operations.V3CollectionsCrudListRequest{
		OrganizationID: censyssdkgointernal.String("<id>"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseEnvelopeListCollectionsResponseV1 != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `List` function may return the following errors:

| Error Type           | Status Code | Content Type             |
| -------------------- | ----------- | ------------------------ |
| sdkerrors.ErrorModel | 401, 403    | application/problem+json |
| sdkerrors.SDKError   | 4XX, 5XX    | \*/\*                    |

### Example

```go
package main

import (
	"context"
	"errors"
	censyssdkgointernal "github.com/censys/censys-sdk-go-internal"
	"github.com/censys/censys-sdk-go-internal/models/operations"
	"github.com/censys/censys-sdk-go-internal/models/sdkerrors"
	"log"
)

func main() {
	ctx := context.Background()

	s := censyssdkgointernal.New(
		censyssdkgointernal.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.Collections.List(ctx, operations.V3CollectionsCrudListRequest{
		OrganizationID: censyssdkgointernal.String("<id>"),
	})
	if err != nil {

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
	censyssdkgointernal "github.com/censys/censys-sdk-go-internal"
	"github.com/censys/censys-sdk-go-internal/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := censyssdkgointernal.New(
		censyssdkgointernal.WithServerURL("https://api.platform.censys.io"),
		censyssdkgointernal.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.Collections.List(ctx, operations.V3CollectionsCrudListRequest{
		OrganizationID: censyssdkgointernal.String("<id>"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseEnvelopeListCollectionsResponseV1 != nil {
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
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
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
	censyssdkgointernal "github.com/censys/censys-sdk-go-internal"
	"github.com/censys/censys-sdk-go-internal/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := censyssdkgointernal.New(
		censyssdkgointernal.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.Collections.List(ctx, operations.V3CollectionsCrudListRequest{
		OrganizationID: censyssdkgointernal.String("<id>"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseEnvelopeListCollectionsResponseV1 != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

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
