# IpmiCapabilitiesAuthStatus


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `AnonymousLoginEnabled`                      | **bool*                                      | :heavy_minus_sign:                           | If true, the server allows anonymous login.  |
| `AuthEachMessage`                            | **bool*                                      | :heavy_minus_sign:                           | If true, each message must be authenticated. |
| `HasAnonymousUsers`                          | **bool*                                      | :heavy_minus_sign:                           | If true, the server has anonymous users.     |
| `HasNamedUsers`                              | **bool*                                      | :heavy_minus_sign:                           | If true, the server supports named users.    |
| `TwoKeyLoginRequired`                        | **bool*                                      | :heavy_minus_sign:                           | The KG field.                                |
| `UserAuthDisabled`                           | **bool*                                      | :heavy_minus_sign:                           | If true, user authentication is disabled.    |