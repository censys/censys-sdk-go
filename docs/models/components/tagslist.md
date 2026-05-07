# TagsList


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `NextPageToken`                                                                     | `*string`                                                                           | :heavy_minus_sign:                                                                  | Token to retrieve the next page of results. Omitted when there are no more results. |
| `Tags`                                                                              | [][components.Tag](../../models/components/tag.md)                                  | :heavy_check_mark:                                                                  | The list of tags.                                                                   |
| `TotalSize`                                                                         | `int64`                                                                             | :heavy_check_mark:                                                                  | Total number of tags visible to the caller in this organization.                    |