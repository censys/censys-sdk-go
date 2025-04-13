# ValidityPeriod


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `LengthSeconds`                                                                 | **int64*                                                                        | :heavy_minus_sign:                                                              | The duration of the certificate's validity period, in seconds.                  |
| `NotAfter`                                                                      | **string*                                                                       | :heavy_minus_sign:                                                              | An RFC-3339-formatted timestamp after which the certificate is no longer valid. |
| `NotBefore`                                                                     | **string*                                                                       | :heavy_minus_sign:                                                              | An RFC-3339-formatted timestamp before which the certificate is not valid.      |