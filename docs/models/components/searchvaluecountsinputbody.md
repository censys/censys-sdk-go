# SearchValueCountsInputBody


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `AndCountConditions`                                                     | [][components.CountCondition](../../models/components/countcondition.md) | :heavy_check_mark:                                                       | Groups of field-value pairs to count matches for.                        |
| `Query`                                                                  | **string*                                                                | :heavy_minus_sign:                                                       | CenQL query string to filter documents                                   |