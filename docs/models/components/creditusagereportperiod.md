# CreditUsageReportPeriod


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `CreditsAdded`                                                 | *int64*                                                        | :heavy_check_mark:                                             | The total amount of credits added during the report period.    |
| `CreditsConsumed`                                              | *int64*                                                        | :heavy_check_mark:                                             | The total amount of credits consumed during the report period. |
| `CreditsExpired`                                               | *int64*                                                        | :heavy_check_mark:                                             | The total amount of credits expired during the report period.  |
| `EndDate`                                                      | [time.Time](https://pkg.go.dev/time#Time)                      | :heavy_check_mark:                                             | The end date of the window for this report period.             |
| `StartDate`                                                    | [time.Time](https://pkg.go.dev/time#Time)                      | :heavy_check_mark:                                             | The start date of the window for this report period.           |
| `TransactionCount`                                             | *int64*                                                        | :heavy_check_mark:                                             | The total number of transactions during the report period.     |