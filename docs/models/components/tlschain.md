# TLSChain


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `FingerprintSha256`                                                          | **string*                                                                    | :heavy_minus_sign:                                                           | SHA 256 fingerprint of the certificate in the certificate chain.             |
| `IssuerDn`                                                                   | **string*                                                                    | :heavy_minus_sign:                                                           | Distinguished name of the entity that has signed and issued the certificate. |
| `SubjectDn`                                                                  | **string*                                                                    | :heavy_minus_sign:                                                           | Distinguished name of the entity that the certificate belongs to.            |