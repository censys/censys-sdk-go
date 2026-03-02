# ServiceOnHostRange


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `EndTime`                                  | [time.Time](https://pkg.go.dev/time#Time)  | :heavy_check_mark:                         | When the service was last observed         |
| `IP`                                       | *string*                                   | :heavy_check_mark:                         | IP address where the service was observed  |
| `Port`                                     | *string*                                   | :heavy_check_mark:                         | Port number where the service was observed |
| `Protocol`                                 | *string*                                   | :heavy_check_mark:                         | Application protocol (e.g., HTTP, HTTPS)   |
| `StartTime`                                | [time.Time](https://pkg.go.dev/time#Time)  | :heavy_check_mark:                         | When the service was first observed        |
| `TransportProtocol`                        | *string*                                   | :heavy_check_mark:                         | Transport protocol (e.g., TCP, UDP, QUIC)  |