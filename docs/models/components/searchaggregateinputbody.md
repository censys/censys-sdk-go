# SearchAggregateInputBody


## Fields

| Field                                    | Type                                     | Required                                 | Description                              |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| `Field`                                  | *string*                                 | :heavy_check_mark:                       | Specify field to aggregate by.           |
| `NumberOfBuckets`                        | *int64*                                  | :heavy_check_mark:                       | Number of buckets to split results into. |
| `Query`                                  | *string*                                 | :heavy_check_mark:                       | CenQL query string to search upon.       |