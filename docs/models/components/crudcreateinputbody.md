# CrudCreateInputBody


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Description`                                                  | **string*                                                      | :heavy_minus_sign:                                             | description of the collection                                  | Hosts with services with AsyncRAT indicator in cert subject DN |
| `Name`                                                         | *string*                                                       | :heavy_check_mark:                                             | name of the collection                                         | Hosts services with AsyncRAT indicator                         |
| `Query`                                                        | *string*                                                       | :heavy_check_mark:                                             | query string to search upon to build the collection            | host.services.cert.parsed.subject_dn: "asyncrat"               |