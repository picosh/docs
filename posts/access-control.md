---
title: Access Control
description: How users can access their services
keywords: [authentication, authorization, access, ssh, certificates]
toc: 2
---

The primary way to access pico services is through SSH keypair authentication. When you create an account, we record the public key you provided which then allows you to access all of our services. This passwordless flow is familiar to most developers and ergonomic. When you add a normal pubkey in the pico.sh TUI, this grants admin access to all services. For a single-user account, this mostly works fine.

# SSH Certificates

We also want to enable secure workflows where people might not want every machine to have admin access to their pico account. Further, we also want to enable teams to use our services without it being a security burden for the account admins.

This is why we also support SSH certificates for authentication and with it we provide more granular access control.

> See our original RFC: https://pico.prose.sh/rfc-007-access-control

All an account admin has to do in order to enable SSH certificates for access control is to add their host CA pubkey to the list of pubkeys authorized to access their account in the pico TUI. After that, they can sign as many pubkeys as they want and with:

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

# admin sends alice-cert.pub to alice and then they can use their keypair
# note: you don't normally need to provide the `-o CertificateFile=` since ssh will find it automatically
# but we wanted to include for completeness
rsync -e "ssh -i ./alice -o CertificateFile=./alice-cert.pub" -rv ./public/ pgs:/site/
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
