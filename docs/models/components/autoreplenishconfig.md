# AutoReplenishConfig


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Amount`                                                                           | **int64*                                                                           | :heavy_minus_sign:                                                                 | The amount of credits to replenish when auto-replenish is triggered.               |
| `Enabled`                                                                          | *bool*                                                                             | :heavy_check_mark:                                                                 | Whether the organization has auto-replenish enabled.                               |
| `Threshold`                                                                        | **int64*                                                                           | :heavy_minus_sign:                                                                 | The threshold at which the organization's credit balance will be auto-replenished. |