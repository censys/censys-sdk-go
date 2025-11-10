# OrganizationMember


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `CreatedAt`                                    | [*time.Time](https://pkg.go.dev/time#Time)     | :heavy_minus_sign:                             | The date and time the user was created.        |
| `Email`                                        | *string*                                       | :heavy_check_mark:                             | The email of the user.                         |
| `FirstLoginTime`                               | [*time.Time](https://pkg.go.dev/time#Time)     | :heavy_minus_sign:                             | The date and time the user first logged in.    |
| `FirstName`                                    | *string*                                       | :heavy_check_mark:                             | The first name of the user.                    |
| `LastName`                                     | *string*                                       | :heavy_check_mark:                             | The last name of the user.                     |
| `LatestLoginTime`                              | [*time.Time](https://pkg.go.dev/time#Time)     | :heavy_minus_sign:                             | The date and time the user last logged in.     |
| `Roles`                                        | []*string*                                     | :heavy_check_mark:                             | The roles this member has in the organization. |
| `UID`                                          | *string*                                       | :heavy_check_mark:                             | The ID of a Censys user.                       |