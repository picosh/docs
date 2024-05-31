---
title: Custom domains
description: A guide to setting up your domain to point to pico services
keywords: [pico, custom, domain]
toc: true
---

All of our services support custom domains and they all work the exact same way.
The way it works is you provide a `CNAME` record and a corresponding `TXT`
record. Then when any of our web services receives traffic from that domain, we
check the `TXT` record to figure out what content to serve the user.

HTTPS will be automatically enabled and a certificate will be retrieved from
Let's Encrypt. In order for this to work, 2 DNS records need to be created:

`CNAME` for the domain to the pico service (subdomains or DNS hosting with CNAME
flattening) or `A` record.

# prose.sh

Setting up a custom domain for prose is the same as all the other services,
except for pgs -- more on that later.

```
CNAME subdomain.yourcustomdomain.com -> prose.sh
```

Resulting in:

```
subdomain.yourcustomdomain.com.         300     IN      CNAME   prose.sh.
```

And a `TXT` record to tell Prose what blog is hosted on that domain at the
subdomain entry `_prose`

```
TXT _prose.subdomain.yourcustomdomain.com -> yourproseusername
```

Resulting in:

```
_prose.subdomain.yourcustomdomain.com.         300     IN      TXT     "hey"
```

Depending on your DNS, this could take some time to fully switch over. We have
an endpoint to check whether or not custom domains are setup:

```
curl -vvvv https://prose.sh/check?domain=xxx
```

# pgs.sh

[pgs.sh](https://pgs.sh) is a little different in that we allow the user to
configure custom domains per project so it's a little different.

And a `TXT` record to tell pgs what project is hosted on that domain at the
subdomain entry `_pgs`.

```
subdomain.yourcustomdomain.com.         300     IN      CNAME   pgs.sh.
_pgs.subdomain.yourcustomdomain.com.    300     IN      TXT
"{user}-{project}"
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

# My DNS does **not** support CNAME flattening

In that case, you need to get the IP address of the service you want to point to
and then use that as an `A` record.

> WARNING: We make **no** guarantees that our IP addresses will stay the same.
> Use at your own risk!
