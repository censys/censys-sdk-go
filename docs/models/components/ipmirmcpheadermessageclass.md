# IpmiRMCPHeaderMessageClass


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Class`                                                         | **int*                                                          | :heavy_minus_sign:                                              | Just the class part of the byte (lower 5 bits of raw)           |
| `IsAck`                                                         | **bool*                                                         | :heavy_minus_sign:                                              | True if the message is an acknowledgment to a previous message. |
| `Name`                                                          | **string*                                                       | :heavy_minus_sign:                                              | The human-readable name of the message class                    |
| `Raw`                                                           | **int*                                                          | :heavy_minus_sign:                                              | The raw message class byte.                                     |