# Ldap


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `AllowsAnonymousBind`                                                  | **bool*                                                                | :heavy_minus_sign:                                                     | Ability to connect with anonymous bind (empty username and password)   |
| `Attributes`                                                           | [][components.LdapAttribute](../../models/components/ldapattribute.md) | :heavy_minus_sign:                                                     | All root DN attributes available via anonymous bind                    |
| `ResultCode`                                                           | **int*                                                                 | :heavy_minus_sign:                                                     | Result or error code returned by LDAP instance upon bind               |