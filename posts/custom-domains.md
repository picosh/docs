---
title: Custom domains
description: A guide to setting up your domain to point to pico services
keywords: [pico, custom, domain]
toc: true
---

All of our services support custom domains and they all work similarly. The way it works is you create two DNS records:

1. A `CNAME` record to send the traffic to us (or an `A`/`AAAA` if you don't have [CNAME flattening](https://developers.cloudflare.com/dns/cname-flattening/))
1. A corresponding `TXT` record which we check to determine what content to serve the user.

The first time we receive traffic on the custom domain, a certificate will be retrieved from Let's Encrypt to enable HTTPS.

Under-the-hood we use Caddy's [on-demand tls](https://caddyserver.com/on-demand-tls) to automatically create certificates for customer domains.

# My DNS does **not** support CNAME flattening

Some DNS providers do not support [CNAME flattening](https://developers.cloudflare.com/dns/cname-flattening/), which means they don't allow you to create a CNAME record for your apex domain, like "example.com".

If you want to use an apex domain without CNAME flattening, you need to get the IP addresses of the service you want to point to and then use them in an `A` record (for IPv4) and an `AAAA` record (for IPv6). You can find the current IP addresses by running `host pgs.sh`.

> WARNING: We make **no** guarantees that our IP addresses will stay the same. Use at your own risk!

```
subdomain.yourcustomdomain.com.         300     IN      A       129.158.37.104.
_pgs.subdomain.yourcustomdomain.com.    300     IN      TXT     user-project
```

# Can I use an `ALIAS` record instead of `CNAME`?

Yes, it should work the exact same way.

```
subdomain.yourcustomdomain.com.         300     IN      ALIAS   pgs.sh.
_pgs.subdomain.yourcustomdomain.com.    300     IN      TXT     user-project
```

# Guides

Here are some guides for known working DNS servers. If you have another DNS server working with pico, please submit a [PR](https://pr.pico.sh/r/erock/pico)!

These guides assume we are setting up `pgs.sh` sites with a root domain.

## Cloudflare

Cloudflare support `CNAME` flattening. [You can read their docs on it](https://developers.cloudflare.com/dns/cname-flattening/).

```
CNAME   @     pgs.sh
TXT     _pgs  user-project
```

## Namecheap

Namecheap supports `CNAME` flattening through a different mechansim by using the `ALIAS` record. [You can read their docs on it](https://www.namecheap.com/support/knowledgebase/article.aspx/10128/2237/how-to-create-an-alias-record/).

```
ALIAS   @     pgs.sh
TXT     _pgs  user-project
```
