---
title: Access Control
description: How users can access their services
keywords: [authentication, authorization, access, ssh, certificates]
toc: 2
---

The primary way to access pico services is through SSH keypair authentication. When you create an account, we record the public key you provided which then allows you to access all of our services. This passwordless flow is familiar to most developers and ergonomic. When you add a normal pubkey in the pico.sh TUI, this grants admin access to all services. For a single-user account, this mostly works fine.

But we also provide the ability to finer-grained access control to your pico account through SSH certificates.

# SSH Certificates

We also want to enable secure workflows where people might not want every machine to have admin access to their pico account. Further, we also want to enable teams to use our services without it being a security burden for the account admins.

This is why we also support SSH certificates for authentication and with it we provide more granular access control.

> See our original RFC: https://pico.prose.sh/rfc-007-access-control

Use cases:

- Machines that only need access to specific services (e.g. CI/CD that only needs pgs access)
- Teams that need shared access to their pico workspace

Setup is simple if you already understand how SSH certificates work. If you need a quick guide, read this blog post [If you’re not using SSH certificates you’re doing SSH wrong](https://smallstep.com/blog/use-ssh-certificates/).

All an account admin needs to do is add their CA pubkey to their Public Keys list. Then the admin requests pubkeys from other teammates or machines, generates a cert pubkey, then sends that cert pubkey back to the user or machine, and finally they use that cert pubkey to authenticate with pico.

Cert pubkeys have additional benefits:

- Access to specific services
- Expiration dates
- Tagged with a readable identifier for access logs

## Workflow

The owner of the pico account **must** generate and manage their own ssh certificate. We will never have access to the host certificate authority (CA) private key.

```bash
# admin creates ssh ca keypair (or uses one they already have)
ssh-keygen -t ed25519 -f ./ca_user_ed25519 -C "pico-ca"

# alice generates a normal ssh keypair (or uses one they already have)
# alice sends pubkey to admin
ssh-keygen -t ed25519 -f alice -C "alice@example.com"

# admin signs alice pubkey with admin ca private key to produce the user certificate
# this generates a new cert-signed pubkey: `alice-cert.pub`
ssh-keygen -s ./ca_user_ed25519 \
           -I "alice@$(date +%F)" \
           -n tuns,pgs \ # this grants alice access to those services
           -V +52w \ # valid for 1 year
           alice.pub

# admin sends alice-cert.pub to alice and then they can use their keypair!

# note: for pico.sh and pipe.pico.sh you need to provide the `-o CertificateFile=`
#  because those services permit non-registered users and your ssh-agent will supply your base keypair
#  first instead of your certified pubkey and those services accept the first key presented.
# we recommend adding these opts to your ssh_config
ssh -o IdentitiesOnly=yes -o CertificateFile=./alice-cert.pub -i ./alice pico.sh

rsync -e "ssh -i ./alice" -rv ./public/ pgs:/site/
```

Only an `admin` in `principals` has full access to pico account management and the pico.sh TUI.

## Principals

- `admin`
- `pgs`
- `tuns`
- `pipe`
- `feeds`
- `prose`
- `pastes`

## Recipes

```bash
ssh-keygen -s ./ca_user_ed25519 -I "alice@$(date +%F)" -n tuns,pgs -V +52w alice.pub
ssh-keygen -s ./ca_user_ed25519 -I "alice@$(date +%F)" -n admin alice.pub
ssh-keygen -s ./ca_user_ed25519 -I "alice@$(date +%F)" -n pipe alice.pub
```

SSH certificates enable finer-grained access control leveraging systems that have existed for years with minimal pico-specific APIs. What's great is pico doesn't need to worry about implementing this functionality, standing on the shoulders of giants. Our only responsibility is to make sure these cert pubkeys can only access the things listed in their principals.
