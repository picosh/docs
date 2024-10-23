---
title: Custom domains
description: A guide to setting up your domain to point to pico services
keywords: [pico, custom, domain]
toc: true
---

All of our services support custom domains and they all work similarly. The way
it works is you create two DNS records:

1. A `CNAME` record to send the traffic to us (or an `A`/`AAAA` if you don't
   have
   [CNAME flattening](https://developers.cloudflare.com/dns/cname-flattening/))
2. A corresponding `TXT` record which we check to determine what content to
   serve the user.

The first time we receive traffic on the custom domain, a certificate will be
retrieved from Let's Encrypt to enable HTTPS.

# prose.sh

Setting up a custom domain for prose is similar to all other services. First you
need to create a CNAME record:

```
CNAME subdomain.yourcustomdomain.com -> prose.sh
```

Resulting in:

```
subdomain.yourcustomdomain.com.   300   IN    CNAME   prose.sh.
```

Next, add a `TXT` record to tell Prose what blog is hosted on that domain at the
subdomain entry `_prose`

```
TXT _prose.subdomain.yourcustomdomain.com -> {user}
```

Where `{user}` is your pico username.

Resulting in:

```
_prose.subdomain.yourcustomdomain.com.    300   IN    TXT   "erock"
```

Depending on your DNS, this could take some time to fully switch over. We have
an endpoint to check whether or not custom domains are setup:

```
curl -i 'https://prose.sh/check?domain=example.com'
```

# pgs.sh

Since [pgs.sh](https://pgs.sh) allows the user to configure custom domains per
project, the `TXT` record format is a little different.

Add a `TXT` record to tell pgs what project is hosted on that domain at the
subdomain entry `_pgs`.

```
subdomain.yourcustomdomain.com.         300     IN      CNAME   pgs.sh.
_pgs.subdomain.yourcustomdomain.com.    300     IN      TXT     "{user}-{project}"
```

## Example: Top-Level Domain

- Custom domain `erock.io`
- User `erock`
- Project `kittens`

Resulting in:

```
erock.io.         300     IN      CNAME   pgs.sh.
_pgs.erock.io.    300     IN      TXT     "erock-kittens"
```

Here's an example of what it looks like inside cloudflare:

![custom domains cloudflare](/custom-domains-cloudflare.png)

## Example: Subdomain

- Custom domain `meow.erock.io`
- User `erock`
- Project `kittens`

Resulting in:

```
meow.erock.io.         300     IN      CNAME   pgs.sh.
_pgs.meow.erock.io.    300     IN      TXT     "erock-kittens"
```

# tuns.sh

Custom domains for [tuns.sh](/tuns) require you to set up `CNAME` and `TXT`
records for the domain/subdomain you would like to use for your forwarded
connection. The `CNAME` record must point to `tuns.sh`. The TXT record name
must be `_sish.customdomain` and contain the SSH key fingerprint used for
creating the tunnel.

You can retrieve your key fingerprint by running:

```
ssh-keygen -lf ~/.ssh/id_rsa | awk '{print $2}'
```

Example:

```
customdomain.example.com.          300     IN      CNAME   tuns.sh.
_sish.customdomain.example.com     300     IN      TXT     "SHA256:mVPwvezndPv/ARoIadVY98vAC0g+P/5633yTC4d/wXE"
```

Once set up, you can then create tunnels via your custom domain like this:

```
ssh -R customdomain.example.com:80:localhost:8000 tuns.sh
```

# My DNS does **not** support CNAME flattening

Some DNS providers do not support
[CNAME flattening](https://developers.cloudflare.com/dns/cname-flattening/),
which means they don't allow you to create a CNAME record for your apex domain,
like "example.com".

If you want to use an apex domain without CNAME flattening, you need to get the
IP addresses of the service you want to point to and then use them in an `A`
record (for IPv4) and an `AAAA` record (for IPv6). You can find the current IP
addresses by running `host pgs.sh`.

> WARNING: We make **no** guarantees that our IP addresses will stay the same.
> Use at your own risk!

```
subdomain.yourcustomdomain.com.         300     IN      A       129.158.37.104.
_pgs.subdomain.yourcustomdomain.com.    300     IN      TXT     "{user}-{project}"
```

# Can I use an `ALIAS` record instead of `CNAME`?

Yes, it should work the exact same way.

```
subdomain.yourcustomdomain.com.         300     IN      ALIAS   pgs.sh.
_pgs.subdomain.yourcustomdomain.com.    300     IN      TXT     "{user}-{project}"
```
