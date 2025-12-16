# UserCredits


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `Balance`                                       | *int64*                                         | :heavy_check_mark:                              | The current credit balance for the user.        |
| `ResetsAt`                                      | [*time.Time](https://pkg.go.dev/time#Time)      | :heavy_minus_sign:                              | The date that the user's credits will be reset. |