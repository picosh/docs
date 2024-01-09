---
title: Getting Started
description: How to use pico services 
keywords: [pico, getting, started]
---

> All of our services use the same database, so when you create an account with
> one, you also have access to all of our other services.

## Create your account with Public-Key Cryptography

We don't need your email address.

To get started, simply ssh into any service's CMS.

```bash
ssh new@prose.sh
```

> Note: `new` is a special username that will always send you to account
> creation, even with multiple accounts associated with your key-pair.

All we need to create an account is your username. This username will be used
for all of your service domains. For example, if your username is `glossy`, we
will create the following domains on your behalf:

```
glossy.prose.sh
glossy.pgs.sh
glossy.pastes.sh
glossy.imgs.sh
glossy.lists.sh
```

After that, just set a username and you're ready to start using our services!
When you SSH again, use your username that you set in the CMS.

```bash
ssh glossy@prose.sh
```
