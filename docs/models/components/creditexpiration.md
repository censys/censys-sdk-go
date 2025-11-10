# CreditExpiration


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `Balance`                                            | *int64*                                              | :heavy_check_mark:                                   | The current balance of the credit expiration.        |
| `CreatedAt`                                          | [*time.Time](https://pkg.go.dev/time#Time)           | :heavy_minus_sign:                                   | The date and time the credit expiration was created. |
| `ExpiresAt`                                          | [*time.Time](https://pkg.go.dev/time#Time)           | :heavy_minus_sign:                                   | The date and time the credit expiration will expire. |