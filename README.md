# Vault Plugin: Solana Auth and Secrets Backend

This repository contains Vault plugin backends for both authentication and secrets engine functionality for Solana wallet keypairs.

## Usage

The binaries and associated SHA256 sums for both plugin backends are prebuilt and uploaded in the repository releases for easy access and downloading. You can alternative build from source for your specific environment as well.

> [!IMPORTANT]
> If the execution environment is MacOs (`darwin`) and using the prebuilt binaries, then you'll likely have to adjust the Apple Quarantine attribute on binaries being used to allow execution because of the lack of Apple codesigning on the builds.
>
> ```bash
> $ xattr -d com.apple.quarantine vault-plugin-<TYPE>-solana
> ```

### Auth

```bash
$ vault plugin register \
    -sha256=$(shasum -a 256 vault-plugin-auth-solana | cut -d ' ' -f1) \
    auth \
    vault-plugin-auth-solana

$ vault auth enable -path=solana vault-plugin-auth-solana
```

### Secrets

```bash
$ vault plugin register \
    -sha256=$(shasum -a 256 vault-plugin-secrets-solana | cut -d ' ' -f1) \
    secret \
    vault-plugin-secrets-solana

$ vault secrets enable -path=solana vault-plugin-secrets-solana
```

## Build Source

The included `Makefile` in the repository contains a target to build the two backend binaries.

```bash
$ make build
```

This will produce the backend binaries at `./buld/plugins/vault-plugin-<TYPE>-solana` to be used with the Vault server.
