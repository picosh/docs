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

Add the CA pubkey itself to your Public Keys list. pico refuses a cert pubkey there: a cert authenticates through the CA that signed it.

Cert pubkeys have additional benefits:

- Access to specific services
- Expiration dates
- Restricted to specific IP addresses
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
- `picd`

A service's name in a cert's principals gives it everything the account can do there. To limit a cert to part of a service, use scopes instead.

## Scopes

A scope is a principal of the form `<service>:<role>:<resource>`:

```bash
ssh-keygen -s ./ca_user_ed25519 -I "ci@$(date +%F)" -n pgs:deploy:blog,pipe:pub:deploys -V +90d ci.pub
```

That cert can deploy the `blog` site on pgs, publish to the `deploys` topic on pipe, and do nothing else. Scopes add up. Leave out the resource to cover everything the account reaches with that role, or end it in `*` to match a prefix.

| Scope | Allows |
|---|---|
| `pgs:read:<project>` | Download files, list them, and open a web tunnel to the site |
| `pgs:deploy:<project>` | Also upload and delete files, and `link`, `unlink`, `rm`, `retain`, `prune` and `cache` |
| `prose:read:<file>`, `pastes:read:<file>`, `feeds:read:<file>` | Download and list |
| `prose:write:<file>`, `pastes:write:<file>`, `feeds:write:<file>` | Also upload, rename and delete, and `rm` and `run` on feeds |
| `pipe:pub:<topic>` | Publish |
| `pipe:sub:<topic>` | Subscribe. Piping needs both roles |
| `tuns:http:<name>` | An HTTP or HTTPS tunnel on that subdomain or custom domain |
| `tuns:sni:<name>` | An SNI tunnel on that name |
| `tuns:tcp:<port>`, `tuns:udp:<port>` | A TCP or UDP tunnel on that port |
| `tuns:alias:<name>` | A TCP alias, as `<name>` or `<name>:<port>` |
| `tuns:connect:<alias>` | Connect to an alias, with the same names |
| `picd:read:<project>`, `picd:deploy:<project>` | See [certificates for one project](https://cd.pico.sh/docs/sharing#certificates-for-one-project), which match names exactly |

- Listing the service itself or `admin` next to a scope gives full access.
- An unknown role or a mistyped scope refuses the cert.
- Certs with scopes can't start admin sessions. They can't change a pgs `acl` or pipe monitors either.
- pgs `link` needs `deploy` on the project and `read` on the one it links to. Deploys that upload a new project and link it need a prefix, such as `pgs:deploy:blog*`.
- pgs `stats` and `forms`, and the pipe `rss` feed, need a scope without a resource.
- pipe topics are named as you type them. Public topics are `public/<topic>`.
- tuns checks the name or port it assigns, so a random one needs a scope that covers it.
- Don't add the signed key itself to your Public Keys list. ssh falls back to the plain key when a cert is refused, and that key has no limits.

## Critical options

pico checks every cert against the CA that signed it, along with its principals and validity window (`-V`).

The only critical option pico supports is `source-address`, which limits the cert to the listed client IP addresses or CIDR ranges:

```bash
ssh-keygen -s ./ca_user_ed25519 -I "ci@$(date +%F)" -n pgs -O source-address=203.0.113.10/32,2001:db8::/32 ci.pub
```

pico rejects a cert with any other critical option, such as `force-command` or `verify-required`. Extensions such as `permit-pty` have no effect.

## Recipes

```bash
ssh-keygen -s ./ca_user_ed25519 -I "alice@$(date +%F)" -n tuns,pgs -V +52w alice.pub
ssh-keygen -s ./ca_user_ed25519 -I "alice@$(date +%F)" -n admin alice.pub
ssh-keygen -s ./ca_user_ed25519 -I "alice@$(date +%F)" -n pipe alice.pub
ssh-keygen -s ./ca_user_ed25519 -I "ci@$(date +%F)" -n pgs -O source-address=203.0.113.10/32 ci.pub
```

SSH certificates enable finer-grained access control leveraging systems that have existed for years with minimal pico-specific APIs. What's great is pico doesn't need to worry about implementing this functionality, standing on the shoulders of giants. Our only responsibility is to make sure these cert pubkeys can only access the things listed in their principals.
