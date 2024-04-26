---
title: Getting started
description: How to use pico services 
keywords: [pico, getting, started]
---

All of our managed services are connected via our main pico.sh SSH service. So
in order to be granted access you will need to first create an account.

# Create your account with Public-Key Cryptography

We don't need your email address, all we need is an SSH keypair. So before you
begin, please make sure you have a valid SSH keypair or
[generate one](/faq#generating-a-new-ssh-key).

To get started, simply use SSH:

```bash
ssh pico.sh
```

Getting `Permission denied (publickey)`?
[Read how to fix it here](faq#permission-denied-when-using-ssh).

![pico-create-account](/pico-create-account.png)

All we need to create an account is your username. This username will be used
for all of your service domains. For example, if your username is `glossy`, we
will create the following domains on your behalf:

```
glossy.prose.sh
glossy.pgs.sh
glossy.pastes.sh
```

After that, just set a username and you're ready to start using our services!

From now on when you SSH to us it'll bring you to our account management system:

```bash
ssh pico.sh
```

# Next Steps

> After you have created your pico account, we **highly** recommend creating an
> [API token](/api-token) and keeping it someplace safe. This is how users can
> recover their account if they lose their SSH private key.

What service did you want to use first? Read the docs to get started with any of
our services:

- [pgs.sh](/pgs) <a href="/plus" class="link-alt-hover">(+)</a>
- [tuns.sh](/tuns) <a href="/plus" class="link-alt-hover">(+)</a>
- [imgs.sh](/imgs) <a href="/plus" class="link-alt-hover">(+)</a>
- [prose.sh](/prose)
- [pastes.sh](/pastes)
- [feeds.sh](/feeds)
