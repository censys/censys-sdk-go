# SearchQueryInputBody


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Fields`                                               | []*string*                                             | :heavy_minus_sign:                                     | specify fields to return in response and ignore others | host.ip                                                |
| `PageSize`                                             | **int64*                                               | :heavy_minus_sign:                                     | amount of results to return per page                   | 1                                                      |
| `PageToken`                                            | **string*                                              | :heavy_minus_sign:                                     | page token for the requested page of search results    |                                                        |
| `Query`                                                | *string*                                               | :heavy_check_mark:                                     | CenQL query string to search upon                      | host.services: (protocol=SSH and not port: 22)         |