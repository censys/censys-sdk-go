# TagsAndComments

## Overview

Endpoints related to asset tagging and commenting

### Available Operations

* [ListComments](#listcomments) - List comments
* [CreateComment](#createcomment) - Create a comment
* [DeleteComment](#deletecomment) - Delete a comment
* [UpdateComment](#updatecomment) - Update a comment
* [ListTags](#listtags) - List tags
* [CreateTag](#createtag) - Create a tag
* [DeleteTag](#deletetag) - Delete a tag
* [GetTag](#gettag) - Get a tag
* [UpdateTag](#updatetag) - Update a tag
* [ListTagAssignments](#listtagassignments) - List tag assignments
* [CreateTagAssignment](#createtagassignment) - Create a tag assignment
* [BulkCreateTagAssignments](#bulkcreatetagassignments) - Bulk create tag assignments
* [BulkDeleteTagAssignments](#bulkdeletetagassignments) - Bulk delete tag assignments
* [DeleteTagAssignment](#deletetagassignment) - Delete a tag assignment
* [ListTagOperations](#listtagoperations) - List tag operations
* [CancelTagOperation](#canceltagoperation) - Cancel a tag operation

## ListComments

Retrieve a paginated list of comments in your organization. Use query parameters to filter by asset, creator, or creation time.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-comments-list-comments" method="get" path="/v3/comments" -->
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

    res, err := s.TagsAndComments.ListComments(ctx, operations.V3CommentsListCommentsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeCommentsList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.V3CommentsListCommentsRequest](../../models/operations/v3commentslistcommentsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V3CommentsListCommentsResponse](../../models/operations/v3commentslistcommentsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 422                      | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## CreateComment

Add a comment on an asset in your organization.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-comments-create-comment" method="post" path="/v3/comments" -->
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

    res, err := s.TagsAndComments.CreateComment(ctx, operations.V3CommentsCreateCommentRequest{
        CreateCommentInputBody: components.CreateCommentInputBody{
            AssetID: "8.8.8.8",
            Body: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeComment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.V3CommentsCreateCommentRequest](../../models/operations/v3commentscreatecommentrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.V3CommentsCreateCommentResponse](../../models/operations/v3commentscreatecommentresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 422                      | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## DeleteComment

Delete a comment. Only the comment's creator or an organization admin can delete a comment. This action is permanent and cannot be undone.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-comments-delete-comment" method="delete" path="/v3/comments/{comment_id}" -->
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

    res, err := s.TagsAndComments.DeleteComment(ctx, operations.V3CommentsDeleteCommentRequest{
        CommentID: "e6faeec1-9f2b-49af-a7e6-119b37d54575",
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
| `request`                                                                                              | [operations.V3CommentsDeleteCommentRequest](../../models/operations/v3commentsdeletecommentrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.V3CommentsDeleteCommentResponse](../../models/operations/v3commentsdeletecommentresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 422                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## UpdateComment

Update the body of an existing comment. Only the comment's creator can update it.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-comments-update-comment" method="put" path="/v3/comments/{comment_id}" -->
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

    res, err := s.TagsAndComments.UpdateComment(ctx, operations.V3CommentsUpdateCommentRequest{
        CommentID: "361c6ff7-2549-4a6d-b212-8ce011a86612",
        UpdateCommentInputBody: components.UpdateCommentInputBody{
            Body: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeComment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.V3CommentsUpdateCommentRequest](../../models/operations/v3commentsupdatecommentrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.V3CommentsUpdateCommentResponse](../../models/operations/v3commentsupdatecommentresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 422                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## ListTags

Retrieve a paginated list of tags in your organization. Private tags created by other users are not included in the results unless your account is an organization admin.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-tags-list-tags" method="get" path="/v3/tags" -->
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

    res, err := s.TagsAndComments.ListTags(ctx, operations.V3TagsListTagsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeTagsList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.V3TagsListTagsRequest](../../models/operations/v3tagslisttagsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.V3TagsListTagsResponse](../../models/operations/v3tagslisttagsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 422                      | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## CreateTag

Create a new tag in your organization. Tags can be used to label and organize assets.<br><br>Specify a privacy setting to control visibility: `private` tags are only visible to you and organization admins, while `shared` tags are visible and manageable by all organization members.<br><br>Tag names must be unique within your organization.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-tags-create-tag" method="post" path="/v3/tags" -->
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

    res, err := s.TagsAndComments.CreateTag(ctx, operations.V3TagsCreateTagRequest{
        CreateTagInputBody: components.CreateTagInputBody{
            Name: "<value>",
            Privacy: components.CreateTagInputBodyPrivacyPrivate,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeTag != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.V3TagsCreateTagRequest](../../models/operations/v3tagscreatetagrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.V3TagsCreateTagResponse](../../models/operations/v3tagscreatetagresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 409, 422                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## DeleteTag

Delete a tag and all of its assignments from your organization. This action is permanent and cannot be undone.<br><br>Only the tag's creator or an organization admin can delete a `private` tag. Tags that are `shared` can be deleted by any organization member.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-tags-delete-tag" method="delete" path="/v3/tags/{tag_id}" -->
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

    res, err := s.TagsAndComments.DeleteTag(ctx, operations.V3TagsDeleteTagRequest{
        TagID: "8e09cd66-475a-4284-88f2-228e9d76dd20",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.V3TagsDeleteTagRequest](../../models/operations/v3tagsdeletetagrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.V3TagsDeleteTagResponse](../../models/operations/v3tagsdeletetagresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 422                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetTag

Retrieve a tag by its ID or name. Tag names are unique within an organization and can be used interchangeably with the tag ID in the path parameter.<br><br>Only tags that are visible to the caller are returned: private tags created by other users are not accessible unless your account is an organization admin.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-tags-get-tag" method="get" path="/v3/tags/{tag_id}" -->
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

    res, err := s.TagsAndComments.GetTag(ctx, operations.V3TagsGetTagRequest{
        TagID: "123e4567-e89b-12d3-a456-426614174000",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeTag != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.V3TagsGetTagRequest](../../models/operations/v3tagsgettagrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.V3TagsGetTagResponse](../../models/operations/v3tagsgettagresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 422                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## UpdateTag

Update an existing tag in your organization. Only the fields provided in the request body will be updated; omitted fields are left unchanged.<br><br>Only the tag's creator or an organization admin can update a `private` tag. Tags with the `shared` setting can be updated by any organization member.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-tags-update-tag" method="put" path="/v3/tags/{tag_id}" -->
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

    res, err := s.TagsAndComments.UpdateTag(ctx, operations.V3TagsUpdateTagRequest{
        TagID: "8367b125-0db2-4688-accc-c2a97a4bdc56",
        UpdateTagInputBody: components.UpdateTagInputBody{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeTag != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.V3TagsUpdateTagRequest](../../models/operations/v3tagsupdatetagrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.V3TagsUpdateTagResponse](../../models/operations/v3tagsupdatetagresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 409, 422            | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## ListTagAssignments

Retrieve a paginated list of assignments for a tag in your organization. Use query parameters to filter results by asset, created_by, or creation time. Only assignments for tags visible to your account are returned.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-tags-list-assignments" method="get" path="/v3/tags/{tag_id}/assignments" -->
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

    res, err := s.TagsAndComments.ListTagAssignments(ctx, operations.V3TagsListAssignmentsRequest{
        TagID: "8b1458f5-044a-4cc5-a600-d602c09ca004",
        AssetID: censyssdkgo.Pointer("8.8.8.8"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeTagAssignmentsList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.V3TagsListAssignmentsRequest](../../models/operations/v3tagslistassignmentsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.V3TagsListAssignmentsResponse](../../models/operations/v3tagslistassignmentsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 422                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## CreateTagAssignment

Assign a tag to an asset. Tag assignments are only visible to members of your organization, depending on the tag's privacy settings. You must have access to the tag to assign it to an asset.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-tags-create-assignment" method="post" path="/v3/tags/{tag_id}/assignments" -->
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

    res, err := s.TagsAndComments.CreateTagAssignment(ctx, operations.V3TagsCreateAssignmentRequest{
        TagID: "2063be9e-6fe8-43db-9f99-815ede3d1b5c",
        CreateTagAssignmentInputBody: components.CreateTagAssignmentInputBody{
            AssetID: "8.8.8.8",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeTagAssignment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.V3TagsCreateAssignmentRequest](../../models/operations/v3tagscreateassignmentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V3TagsCreateAssignmentResponse](../../models/operations/v3tagscreateassignmentresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 409, 422            | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## BulkCreateTagAssignments

Start a long-running operation that assigns a tag to every asset matching a CenQL query. The operation runs asynchronously; use the returned operation to track its progress.<br><br>A tag can hold a limited number of asset assignments, and the limit depends on your plan. If the operation reaches that limit before every matching asset is tagged, it finishes with status `limit_reached`.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-tags-bulk-create-assignments" method="post" path="/v3/tags/{tag_id}/assignments/bulk-create" -->
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

    res, err := s.TagsAndComments.BulkCreateTagAssignments(ctx, operations.V3TagsBulkCreateAssignmentsRequest{
        TagID: "1de943e6-02d8-47bb-8655-559621fc968c",
        BulkCreateTagAssignmentsInputBody: components.BulkCreateTagAssignmentsInputBody{
            Query: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeTagOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.V3TagsBulkCreateAssignmentsRequest](../../models/operations/v3tagsbulkcreateassignmentsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.V3TagsBulkCreateAssignmentsResponse](../../models/operations/v3tagsbulkcreateassignmentsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 409, 422            | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## BulkDeleteTagAssignments

Start a long-running operation that removes a tag from its assigned assets, optionally restricted to a creation-time window. The operation runs asynchronously; use the returned operation to track its progress.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-tags-bulk-delete-assignments" method="post" path="/v3/tags/{tag_id}/assignments/bulk-delete" -->
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

    res, err := s.TagsAndComments.BulkDeleteTagAssignments(ctx, operations.V3TagsBulkDeleteAssignmentsRequest{
        TagID: "4fb0c28f-5f77-4d94-8d15-c0c621009061",
        BulkDeleteTagAssignmentsInputBody: components.BulkDeleteTagAssignmentsInputBody{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeTagOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.V3TagsBulkDeleteAssignmentsRequest](../../models/operations/v3tagsbulkdeleteassignmentsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.V3TagsBulkDeleteAssignmentsResponse](../../models/operations/v3tagsbulkdeleteassignmentsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 409, 422            | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## DeleteTagAssignment

Remove a tag assignment from an asset. This action is permanent and cannot be undone. Removing an assignment only detaches the tag from the specified asset; the tag itself is not deleted. Only the tag's creator or an organization admin can delete an assignment for a `private` tag. Assignments for `shared` tags can be deleted by any organization member.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-tags-delete-assignment" method="delete" path="/v3/tags/{tag_id}/assignments/{assignment_id}" -->
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

    res, err := s.TagsAndComments.DeleteTagAssignment(ctx, operations.V3TagsDeleteAssignmentRequest{
        TagID: "ad98c1dc-289b-4487-b11f-d41cd4391806",
        AssignmentID: "35060ce7-9fe8-4c5f-9889-efb0572473c2",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.V3TagsDeleteAssignmentRequest](../../models/operations/v3tagsdeleteassignmentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V3TagsDeleteAssignmentResponse](../../models/operations/v3tagsdeleteassignmentresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 422                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## ListTagOperations

Retrieve a paginated list of bulk tag operations. Provide a tag ID in the path to scope the listing to a single tag, or `-` to list operations across all tags in your organization. Use query parameters to filter by type or status.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-tags-list-operations" method="get" path="/v3/tags/{tag_id}/operations" -->
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

    res, err := s.TagsAndComments.ListTagOperations(ctx, operations.V3TagsListOperationsRequest{
        TagID: "123e4567-e89b-12d3-a456-426614174000",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeTagOperationsList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.V3TagsListOperationsRequest](../../models/operations/v3tagslistoperationsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.V3TagsListOperationsResponse](../../models/operations/v3tagslistoperationsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 409, 422                 | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## CancelTagOperation

Request cancellation of an in-progress bulk tag operation. Cancellation is cooperative: the operation finishes the page it is currently processing, then stops. Work already committed before cancellation is preserved.<br><br>Cancellation is not a rollback — it stops further work but does not remove assignments the operation already created. To remove assignments, use the bulk-delete endpoint. An operation that has already finished cannot be cancelled.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-tags-cancel-operation" method="post" path="/v3/tags/{tag_id}/operations/{operation_id}/cancel" -->
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

    res, err := s.TagsAndComments.CancelTagOperation(ctx, operations.V3TagsCancelOperationRequest{
        TagID: "59da45bf-13dd-4601-8b86-c3af7e33fd9a",
        OperationID: "174d4e39-b9c5-4b3f-aa69-5937abd03e40",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeTagOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.V3TagsCancelOperationRequest](../../models/operations/v3tagscanceloperationrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.V3TagsCancelOperationResponse](../../models/operations/v3tagscanceloperationresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 409, 422            | application/problem+json      |
| sdkerrors.ErrorModel          | 500                           | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |