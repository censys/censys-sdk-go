# TagAssignmentsList


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Assignments`                                                                       | [][components.TagAssignment](../../models/components/tagassignment.md)              | :heavy_check_mark:                                                                  | The list of tag assignments.                                                        |
| `NextPageToken`                                                                     | `*string`                                                                           | :heavy_minus_sign:                                                                  | Token to retrieve the next page of results. Omitted when there are no more results. |
| `TotalSize`                                                                         | `int64`                                                                             | :heavy_check_mark:                                                                  | Total number of assignments matching the filters.                                   |