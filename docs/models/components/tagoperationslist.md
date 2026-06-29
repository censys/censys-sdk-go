# TagOperationsList


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `NextPageToken`                                                                     | `*string`                                                                           | :heavy_minus_sign:                                                                  | Token to retrieve the next page of results. Omitted when there are no more results. |
| `Operations`                                                                        | [][components.TagOperation](../../models/components/tagoperation.md)                | :heavy_check_mark:                                                                  | The list of tag operations.                                                         |
| `TotalSize`                                                                         | `int64`                                                                             | :heavy_check_mark:                                                                  | Total number of operations matching the filters.                                    |