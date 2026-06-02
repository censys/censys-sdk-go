# CommentsList


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Comments`                                                                          | [][components.Comment](../../models/components/comment.md)                          | :heavy_check_mark:                                                                  | The list of comments.                                                               |
| `NextPageToken`                                                                     | `*string`                                                                           | :heavy_minus_sign:                                                                  | Token to retrieve the next page of results. Omitted when there are no more results. |
| `TotalSize`                                                                         | `int64`                                                                             | :heavy_check_mark:                                                                  | Total number of comments matching the filters.                                      |