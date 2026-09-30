# V3ThreathuntingInvestigationsJobsResultsResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `ResponseStream`                                                   | `io.ReadCloser`                                                    | :heavy_minus_sign:                                                 | Investigation results archive                                      |
| `Headers`                                                          | map[string][]`string`                                              | :heavy_check_mark:                                                 | N/A                                                                |