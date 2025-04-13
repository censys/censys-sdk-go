# Amqp


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ExplicitTLS`                                                       | **bool*                                                             | :heavy_minus_sign:                                                  | Connected via a TLS connection after initial handshake              |
| `ImplicitTLS`                                                       | **bool*                                                             | :heavy_minus_sign:                                                  | Connected via a TLS wrapped connection (AMQPS)                      |
| `ProtocolID`                                                        | [*components.AmqpProtocol](../../models/components/amqpprotocol.md) | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Version`                                                           | [*components.AmqpVersion](../../models/components/amqpversion.md)   | :heavy_minus_sign:                                                  | N/A                                                                 |