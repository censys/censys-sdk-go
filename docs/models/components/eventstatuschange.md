# EventStatusChange


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `EventTime`                                                   | [time.Time](https://pkg.go.dev/time#Time)                     | :heavy_check_mark:                                            | N/A                                                           |
| `NewStatus`                                                   | [components.NewStatus](../../models/components/newstatus.md)  | :heavy_check_mark:                                            | N/A                                                           |
| `OldStatus`                                                   | [*components.OldStatus](../../models/components/oldstatus.md) | :heavy_minus_sign:                                            | N/A                                                           |
| `Reason`                                                      | **string*                                                     | :heavy_minus_sign:                                            | N/A                                                           |