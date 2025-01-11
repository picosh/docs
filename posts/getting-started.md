---
title: Getting started
description: How to use pico services
keywords: [pico, getting, started]
---

All of our managed services are connected by our main pico.sh SSH service. So in
order to be granted access you will need to create an account.

# Create your account with Public-Key Cryptography

All we need is an SSH keypair. So before you begin, please make sure you have a
valid SSH keypair or [generate one](/faq#how-do-i-generate-an-ssh-key).

> By signing up to pico, you agree to our [privacy policy](/privacy) and our
> [terms of service](/ops).

To get started, simply use SSH:

```bash
ssh pico.sh
```

> Getting an error? Read our FAQ,
> [why can't I login to pico?](faq#why-cant-i-login-to-pico).

All we need to create an account is your username. This username will be used
for all of your service domains. For example, if your username is `glossy`, we
will create the following domains on your behalf:

```
glossy.prose.sh
glossy.pgs.sh
glossy.pastes.sh
```

From now on when you SSH it'll bring you to our account management system:

```bash
ssh pico.sh
```

# Next Steps

> After you have created your pico account, we **highly** recommend creating an
> [API token](/api-token) and keeping it someplace safe. This is how users can
> recover their account if they lose their SSH private key.

What service did you want to use first? Read the docs to get started with any of
our services:

- [tuns](/tuns) <a href="/plus" class="link-alt-hover">(+)</a>
- [pages](/pgs) <a href="/plus" class="link-alt-hover">(+)</a>
- [pipe](https://pipe.pico.sh)
- [prose](/prose)
- [pastes](/pastes)
- [rss-to-email](/feeds)
