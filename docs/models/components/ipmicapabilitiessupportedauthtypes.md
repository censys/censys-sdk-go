# IpmiCapabilitiesSupportedAuthTypes


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Extended`                                        | **bool*                                           | :heavy_minus_sign:                                | If true, the extended capabilities are present.   |
| `Md2`                                             | **bool*                                           | :heavy_minus_sign:                                | True if the MD2 AuthType is supported.            |
| `Md5`                                             | **bool*                                           | :heavy_minus_sign:                                | True if the MD5 AuthType is supported.            |
| `None`                                            | **bool*                                           | :heavy_minus_sign:                                | True if the None AuthType is supported.           |
| `OemProprietary`                                  | **bool*                                           | :heavy_minus_sign:                                | True if the OEM Proprietary AuthType is supported |
| `Password`                                        | **bool*                                           | :heavy_minus_sign:                                | True if the Password AuthType is supported.       |
| `Raw`                                             | **int*                                            | :heavy_minus_sign:                                | The raw byte, with the bit mask etc               |