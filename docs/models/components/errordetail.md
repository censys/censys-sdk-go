# ErrorDetail


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Location`                                                             | **string*                                                              | :heavy_minus_sign:                                                     | Where the error occurred, e.g. 'body.items[3].tags' or 'path.thing-id' |
| `Message`                                                              | **string*                                                              | :heavy_minus_sign:                                                     | Error message text                                                     |
| `Value`                                                                | *any*                                                                  | :heavy_minus_sign:                                                     | The value at the given location                                        |