# SearchAggregateInputBody


## Fields

| Field                                   | Type                                    | Required                                | Description                             | Example                                 |
| --------------------------------------- | --------------------------------------- | --------------------------------------- | --------------------------------------- | --------------------------------------- |
| `Field`                                 | *string*                                | :heavy_check_mark:                      | field to aggregate by                   | web.endpoints.http.html_title           |
| `NumberOfBuckets`                       | *int64*                                 | :heavy_check_mark:                      | number of buckets to split results into | 100                                     |
| `Query`                                 | *string*                                | :heavy_check_mark:                      | CenQL query string to search upon       | web: *                                  |