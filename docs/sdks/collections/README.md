# Collections
(*Collections*)

## Overview

Endpoints related to the Collections product

### Available Operations

* [List](#list) - List Collections
* [Create](#create) - Create a Collection
* [Delete](#delete) - Delete a Collection
* [Get](#get) - Get a Collection
* [Update](#update) - Update a Collection
* [ListEvents](#listevents) - List a Collection's events
* [Aggregate](#aggregate) - Search / Aggregate
* [Search](#search) - Search / Query

## List

List Collections

### Example Usage

```go
package main

import(
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

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.V3CollectionsCrudListRequest](../../models/operations/v3collectionscrudlistrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.V3CollectionsCrudListResponse](../../models/operations/v3collectionscrudlistresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## Create

Create a Collection

### Example Usage

```go
package main

import(
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

    res, err := s.Collections.Create(ctx, operations.V3CollectionsCrudCreateRequest{
        OrganizationID: censyssdkgointernal.String("<id>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.V3CollectionsCrudCreateRequest](../../models/operations/v3collectionscrudcreaterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.V3CollectionsCrudCreateResponse](../../models/operations/v3collectionscrudcreateresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## Delete

Delete a Collection

### Example Usage

```go
package main

import(
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

    res, err := s.Collections.Delete(ctx, operations.V3CollectionsCrudDeleteRequest{
        OrganizationID: censyssdkgointernal.String("<id>"),
        CollectionUID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.V3CollectionsCrudDeleteRequest](../../models/operations/v3collectionscruddeleterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.V3CollectionsCrudDeleteResponse](../../models/operations/v3collectionscruddeleteresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## Get

Get a Collection

### Example Usage

```go
package main

import(
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

    res, err := s.Collections.Get(ctx, operations.V3CollectionsCrudGetRequest{
        OrganizationID: censyssdkgointernal.String("<id>"),
        CollectionUID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.V3CollectionsCrudGetRequest](../../models/operations/v3collectionscrudgetrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.V3CollectionsCrudGetResponse](../../models/operations/v3collectionscrudgetresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## Update

Update a Collection

### Example Usage

```go
package main

import(
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

    res, err := s.Collections.Update(ctx, operations.V3CollectionsCrudUpdateRequest{
        OrganizationID: censyssdkgointernal.String("<id>"),
        CollectionUID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.V3CollectionsCrudUpdateRequest](../../models/operations/v3collectionscrudupdaterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.V3CollectionsCrudUpdateResponse](../../models/operations/v3collectionscrudupdateresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## ListEvents

List a Collection's events

### Example Usage

```go
package main

import(
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

    res, err := s.Collections.ListEvents(ctx, operations.V3CollectionsListEventsRequest{
        OrganizationID: censyssdkgointernal.String("<id>"),
        CollectionUID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeCollectionEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.V3CollectionsListEventsRequest](../../models/operations/v3collectionslisteventsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.V3CollectionsListEventsResponse](../../models/operations/v3collectionslisteventsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## Aggregate

Run an aggregation via a Collection data set

### Example Usage

```go
package main

import(
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

    res, err := s.Collections.Aggregate(ctx, operations.V3CollectionsSearchAggregateRequest{
        OrganizationID: censyssdkgointernal.String("<id>"),
        CollectionUID: "<id>",
        SearchAggregateInputBody: components.SearchAggregateInputBody{
            Field: "<value>",
            NumberOfBuckets: 590414,
            Query: "<value>",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.V3CollectionsSearchAggregateRequest](../../models/operations/v3collectionssearchaggregaterequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.V3CollectionsSearchAggregateResponse](../../models/operations/v3collectionssearchaggregateresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## Search

Run a query via a Collection data set

### Example Usage

```go
package main

import(
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

    res, err := s.Collections.Search(ctx, operations.V3CollectionsSearchQueryRequest{
        OrganizationID: censyssdkgointernal.String("<id>"),
        CollectionUID: "<id>",
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

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.V3CollectionsSearchQueryRequest](../../models/operations/v3collectionssearchqueryrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.V3CollectionsSearchQueryResponse](../../models/operations/v3collectionssearchqueryresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorModel     | 401, 403                 | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |