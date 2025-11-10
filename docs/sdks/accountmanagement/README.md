# AccountManagement
(*AccountManagement*)

## Overview

Endpoints related to the Account Management product

### Available Operations

* [GetOrganizationDetails](#getorganizationdetails) - Get organization details
* [GetOrganizationCredits](#getorganizationcredits) - Get organization credit statistics
* [GetOrganizationCreditUsage](#getorganizationcreditusage) - Get organization credit usage
* [InviteUserToOrganization](#inviteusertoorganization) - Invite user to organization
* [ListOrganizationMembers](#listorganizationmembers) - List organization members
* [RemoveOrganizationMember](#removeorganizationmember) - Remove member from organization
* [UpdateOrganizationMember](#updateorganizationmember) - Update a member's roles in an organization
* [GetMemberCreditUsage](#getmembercreditusage) - Get member credit usage

## GetOrganizationDetails

Retrieve an organization's details, including the count of organization members broken down by role and organization settings such as AI training and MFA requirements.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-accountmanagement-org-details" method="get" path="/v3/accounts/organizations/{organization_id}" -->
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

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.V3AccountmanagementOrgDetailsRequest](../../models/operations/v3accountmanagementorgdetailsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.V3AccountmanagementOrgDetailsResponse](../../models/operations/v3accountmanagementorgdetailsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 422                 | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetOrganizationCredits

Retrieve credit balance and expiration information for an organization. <br><br>Credits expire 12 months after they are acquired.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-accountmanagement-org-credits" method="get" path="/v3/accounts/organizations/{organization_id}/credits" -->
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
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AccountManagement.GetOrganizationCredits(ctx, operations.V3AccountmanagementOrgCreditsRequest{
        OrganizationID: "11111111-2222-3333-4444-555555555555",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeOrganizationCredits != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.V3AccountmanagementOrgCreditsRequest](../../models/operations/v3accountmanagementorgcreditsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.V3AccountmanagementOrgCreditsResponse](../../models/operations/v3accountmanagementorgcreditsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 422                 | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetOrganizationCreditUsage

Retrieve credit consumption information for an organization for a specific day.<br><br>Admins can obtain credit usage information for all users in their organization. Members may only retrieve usage information for their own account.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-accountmanagement-org-credits-usage" method="get" path="/v3/accounts/organizations/{organization_id}/credits/usage" -->
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
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AccountManagement.GetOrganizationCreditUsage(ctx, operations.V3AccountmanagementOrgCreditsUsageRequest{
        OrganizationID: "11111111-2222-3333-4444-555555555555",
        Date: "2025-11-01",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeCreditUsageReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.V3AccountmanagementOrgCreditsUsageRequest](../../models/operations/v3accountmanagementorgcreditsusagerequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.V3AccountmanagementOrgCreditsUsageResponse](../../models/operations/v3accountmanagementorgcreditsusageresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 400, 403, 404, 422            | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## InviteUserToOrganization

Invite a user to an organization. The user will receive an email to join the organization. This is equivalent to [adding a new member via the UI](https://docs.censys.com/docs/platform-org-management#invite-members).<br><br>Only users with the Admin role in the provided organization can perform this operation.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-accountmanagement-invite-user-to-org" method="post" path="/v3/accounts/organizations/{organization_id}/invitations" -->
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
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AccountManagement.InviteUserToOrganization(ctx, operations.V3AccountmanagementInviteUserToOrgRequest{
        OrganizationID: "11111111-2222-3333-4444-555555555555",
        InviteMemberInputBody: components.InviteMemberInputBody{
            Email: "user@example.com",
        },
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.V3AccountmanagementInviteUserToOrgRequest](../../models/operations/v3accountmanagementinviteusertoorgrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.V3AccountmanagementInviteUserToOrgResponse](../../models/operations/v3accountmanagementinviteusertoorgresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 422                 | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## ListOrganizationMembers

Retrieve a paginated list of an organization's members and their user details, including their user ID, email, name, creation time, and roles.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-accountmanagement-list-org-members" method="get" path="/v3/accounts/organizations/{organization_id}/members" -->
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
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AccountManagement.ListOrganizationMembers(ctx, operations.V3AccountmanagementListOrgMembersRequest{
        OrganizationID: "11111111-2222-3333-4444-555555555555",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeOrganizationMembersList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.V3AccountmanagementListOrgMembersRequest](../../models/operations/v3accountmanagementlistorgmembersrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.V3AccountmanagementListOrgMembersResponse](../../models/operations/v3accountmanagementlistorgmembersresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404                      | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## RemoveOrganizationMember

Remove a user from an organization. This is equivalent to [removing a member via the UI](https://docs.censys.com/docs/platform-org-management#remove-members).<br><br>Only users with the Admin role in the provided organization can perform this operation.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-accountmanagement-remove-org-member" method="delete" path="/v3/accounts/organizations/{organization_id}/members/{user_id}" -->
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
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AccountManagement.RemoveOrganizationMember(ctx, operations.V3AccountmanagementRemoveOrgMemberRequest{
        OrganizationID: "11111111-2222-3333-4444-555555555555",
        UserID: "11111111-2222-3333-4444-555555555555",
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.V3AccountmanagementRemoveOrgMemberRequest](../../models/operations/v3accountmanagementremoveorgmemberrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.V3AccountmanagementRemoveOrgMemberResponse](../../models/operations/v3accountmanagementremoveorgmemberresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 403, 404, 409, 422            | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## UpdateOrganizationMember

Update the roles assigned to an organization member. This operation replaces a member's roles with the list provided in the request body. To remove all roles from a member, provide an empty list. To completely remove a member from an organization, use the [remove member endpoint](https://docs.censys.com/reference/v3-accountmanagement-remove-org-member).<br><br>Only users with the Admin role in the provided organization can perform this operation.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-accountmanagement-update-org-member" method="patch" path="/v3/accounts/organizations/{organization_id}/members/{user_id}" -->
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
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AccountManagement.UpdateOrganizationMember(ctx, operations.V3AccountmanagementUpdateOrgMemberRequest{
        OrganizationID: "11111111-2222-3333-4444-555555555555",
        UserID: "11111111-2222-3333-4444-555555555555",
        UpdateMemberRoleInputBody: components.UpdateMemberRoleInputBody{
            Roles: nil,
        },
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.V3AccountmanagementUpdateOrgMemberRequest](../../models/operations/v3accountmanagementupdateorgmemberrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.V3AccountmanagementUpdateOrgMemberResponse](../../models/operations/v3accountmanagementupdateorgmemberresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 400, 403, 404, 422            | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetMemberCreditUsage

Retrieve credit consumption information for an organization member for a specific day.<br><br>This endpoint does not cost any credits to execute.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3-accountmanagement-member-credits-usage" method="get" path="/v3/accounts/organizations/{organization_id}/members/{user_id}/credits/usage" -->
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
        censyssdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.AccountManagement.GetMemberCreditUsage(ctx, operations.V3AccountmanagementMemberCreditsUsageRequest{
        OrganizationID: "11111111-2222-3333-4444-555555555555",
        UserID: "11111111-2222-3333-4444-555555555555",
        Date: "2025-11-01",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseEnvelopeCreditUsageReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.V3AccountmanagementMemberCreditsUsageRequest](../../models/operations/v3accountmanagementmembercreditsusagerequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.V3AccountmanagementMemberCreditsUsageResponse](../../models/operations/v3accountmanagementmembercreditsusageresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.AuthenticationError | 401                           | application/json              |
| sdkerrors.ErrorModel          | 400, 403, 404, 422            | application/problem+json      |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |