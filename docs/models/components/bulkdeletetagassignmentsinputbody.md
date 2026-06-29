# BulkDeleteTagAssignmentsInputBody


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `CreatedAfter`                                                       | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | RFC3339 timestamp. Only delete assignments created after this time.  |
| `CreatedBefore`                                                      | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | RFC3339 timestamp. Only delete assignments created before this time. |