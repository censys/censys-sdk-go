# ServicesOnHostResponse


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `NextPageToken`                                                                  | `string`                                                                         | :heavy_check_mark:                                                               | A token that can be used to retrieve the next page of ranges.                    |
| `Ranges`                                                                         | [][components.ServiceOnHostRange](../../models/components/serviceonhostrange.md) | :heavy_check_mark:                                                               | The list of requested services.                                                  |