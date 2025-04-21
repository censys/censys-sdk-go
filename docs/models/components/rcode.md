# RCode

A enumerated field indicating the result of the request. The most common values are defined in RFC 1035.


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `RCodeUnknown`        |                       |
| `RCodeSuccess`        | success               |
| `RCodeFormatError`    | format_error          |
| `RCodeServerFailure`  | server_failure        |
| `RCodeNameError`      | name_error            |
| `RCodeNotImplemented` | not_implemented       |
| `RCodeRefused`        | refused               |
| `RCodeYxDomain`       | yx_domain             |
| `RCodeYxRrset`        | yx_rrset              |
| `RCodeNxRrset`        | nx_rrset              |
| `RCodeNotAuth`        | not_auth              |
| `RCodeNotZone`        | not_zone              |
| `RCodeBadSig`         | bad_sig               |
| `RCodeBadKey`         | bad_key               |
| `RCodeBadTime`        | bad_time              |
| `RCodeBadMode`        | bad_mode              |
| `RCodeBadName`        | bad_name              |
| `RCodeBadAlg`         | bad_alg               |
| `RCodeBadTrunc`       | bad_trunc             |
| `RCodeBadCookie`      | bad_cookie            |