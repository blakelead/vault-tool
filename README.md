# Vault Tool

A command line tool that makes working with Hashicorp Vault easier.

## Disclaimer

This tool is in early stage of development and should not be used to handle important/production secrets.

## Usage

```bash
Usage:
  vault-tool [command]

Available Commands:
  delete      Delete secrets
  dump        Dump secrets to stdout
  help        Help about any command
  migrate     Migrate secrets

Flags:
      --config string   config file (default is $HOME/.vault-tool.yaml) (default ".vault-tool.yaml")
  -h, --help            help for vault-tool
  -v, --version         version for vault-tool
```

Operations apply to kv1 and kv2 secrets.

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

Print all secrets under a path:

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

Copy secrets from one Vault to another (or from one path to another in the same Vault):

```bash
> vault-tool migrate --config vault-tool.yaml secret/path secret/otherpath
```

Delete all secrets under a path:

```bash
> vault-tool delete --config vault-tool.yaml secret/path
```

## Planned Features

- Add tests
- Improve configuration
- Mask default vault env variables
- Add other types of authentication (userpass, ldap, certs)
- Check token capabilities/ttl before write operation
- Run all operations concurrently
- Write secrets from JSON dump
- Create env variable from secret
- Add regex capabilities in path
- Autocompletion