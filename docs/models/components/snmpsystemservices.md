# SnmpSystemServices


## Fields

| Field                              | Type                               | Required                           | Description                        |
| ---------------------------------- | ---------------------------------- | ---------------------------------- | ---------------------------------- |
| `Layer1`                           | **bool*                            | :heavy_minus_sign:                 | Physical (e.g. repeaters)          |
| `Layer2`                           | **bool*                            | :heavy_minus_sign:                 | Datalink/subnetwork (e.g. bridges) |
| `Layer3`                           | **bool*                            | :heavy_minus_sign:                 | Internet (e.g. IP gateways)        |
| `Layer4`                           | **bool*                            | :heavy_minus_sign:                 | End-to-end (e.g. IP hosts)         |
| `Layer5`                           | **bool*                            | :heavy_minus_sign:                 | OSI layer 5                        |
| `Layer6`                           | **bool*                            | :heavy_minus_sign:                 | OSI layer 6                        |
| `Layer7`                           | **bool*                            | :heavy_minus_sign:                 | Applications (e.g. mail relays)    |