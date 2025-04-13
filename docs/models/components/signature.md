# Signature


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `SelfSigned`                                                        | **bool*                                                             | :heavy_minus_sign:                                                  | Whether the certificate was signed by its own key.                  |
| `SignatureAlgorithm`                                                | [*components.KeyAlgorithm](../../models/components/keyalgorithm.md) | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Valid`                                                             | **bool*                                                             | :heavy_minus_sign:                                                  | Whether the signature is valid.                                     |
| `Value`                                                             | **string*                                                           | :heavy_minus_sign:                                                  | Contents of the signature.                                          |