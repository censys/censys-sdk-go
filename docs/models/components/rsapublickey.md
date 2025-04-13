# RsaPublicKey


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Exponent`                                        | **int64*                                          | :heavy_minus_sign:                                | The RSA key's public exponent (e).                |
| `Length`                                          | **int64*                                          | :heavy_minus_sign:                                | Bit-length of the RSA modulus.                    |
| `Modulus`                                         | **string*                                         | :heavy_minus_sign:                                | The RSA key's modulus (n) in big-endian encoding. |