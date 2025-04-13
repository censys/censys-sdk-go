# SearchQueryInputBody


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Fields`                                                | []*string*                                              | :heavy_minus_sign:                                      | Specify fields to return in response and ignore others. |
| `PageSize`                                              | **int64*                                                | :heavy_minus_sign:                                      | Amount of results to return per page.                   |
| `PageToken`                                             | **string*                                               | :heavy_minus_sign:                                      | Page token for the requested page of search results.    |
| `Query`                                                 | *string*                                                | :heavy_check_mark:                                      | CenQL query string to search upon.                      |