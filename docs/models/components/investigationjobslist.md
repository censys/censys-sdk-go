# InvestigationJobsList


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Jobs`                                                                                     | [][components.InvestigationJob](../../models/components/investigationjob.md)               | :heavy_check_mark:                                                                         | The caller's investigations in this organization, newest first.                            |
| `NextPageToken`                                                                            | `*string`                                                                                  | :heavy_minus_sign:                                                                         | Token to retrieve the next page of investigations. Omitted when there are no more results. |