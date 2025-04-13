# Vnc


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ConnectionFailedReason`                                           | **string*                                                          | :heavy_minus_sign:                                                 | If server terminates handshake, the reason offered (if any)        |
| `DesktopName`                                                      | **string*                                                          | :heavy_minus_sign:                                                 | Desktop name provided by the server, capped at 255 bytes           |
| `PixelEncoding`                                                    | [*components.VncKeyValue](../../models/components/vnckeyvalue.md)  | :heavy_minus_sign:                                                 | N/A                                                                |
| `ScreenInfo`                                                       | [*components.DesktopInfo](../../models/components/desktopinfo.md)  | :heavy_minus_sign:                                                 | N/A                                                                |
| `SecurityTypes`                                                    | [][components.VncKeyValue](../../models/components/vnckeyvalue.md) | :heavy_minus_sign:                                                 | server-specified security options                                  |
| `Version`                                                          | **string*                                                          | :heavy_minus_sign:                                                 | N/A                                                                |