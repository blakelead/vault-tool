# Vault Tool

A command line tool that makes working with Hashicorp Vault easier.

## Disclaimer

This tool is in early stage of development and should not be used to handle production Vault secrets.

## Quick start

Create a configuration file:

```yaml
# vault-tool.yaml
source:
  address: https://vault-server.com # Address of Vault server
  token: s.token1                   # Token with enough rights to perform wanted tasks
  insecure: true                    # Skip TLS verification (default: false)
  readonly: true                    # Prevent write/delete operations (default: false)

destination:
  address: https://other-vault-server.com
  token: s.token2
  insecure: true
```

Print secrets in path recursively:

```bash
> vault-tool dump --config vault-tool.yaml secret/path
{
    "secret/path/subpath/secret1": {
      "key1": "value1"
    },
    "secret/path/subpath/secret2": {
      "key2": "value2"
    },
}
```

Copy secrets from one Vault to another:

```bash
> vault-tool migrate --config vault-tool.yaml secret/path secret/otherpath
```

Delete entire path:

```bash
> vault-tool delete --config vault-tool.yaml secret/path
```

## Implemented Features

- Copy secrets from one path to another in the same or different Vaults
- Output secrets in stdout in JSON format
- Delete secret or path

## Planned Features

- Add tests
- Improve configuration
- Mask default vault env variables
- Add other types of authentication (userpass, ldap, certs)
- Check token capabilities/ttl before write operation
- Run all operations concurrently
- Write secrets from JSON
- Create env variable from secret
- Add regex capabilities in path
- Autocompletion