# HostAssetWithMatchedServices


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Extensions`                                                             | map[string]*any*                                                         | :heavy_check_mark:                                                       | N/A                                                                      |
| `MatchedServices`                                                        | [][components.MatchedService](../../models/components/matchedservice.md) | :heavy_minus_sign:                                                       | The host services that match the query.                                  |
| `Resource`                                                               | [components.Host](../../models/components/host.md)                       | :heavy_check_mark:                                                       | N/A                                                                      |