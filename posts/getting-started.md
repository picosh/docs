---
title: Getting started
description: How to use pico services
keywords: [pico, getting, started]
---

Access to our services requires creating an account using our pico.sh TUI.

![pico.sh tui](/tui.png)

# Create your account with Public-Key Cryptography

All we need is an SSH keypair. So before you begin, please make sure you have a
valid SSH keypair or [generate one](/faq#how-do-i-generate-an-ssh-key).

> By signing up to pico, you agree to our [privacy policy](/privacy) and our
> [terms of service](/ops).

To get started, `ssh` into our TUI console:

```bash
ssh pico.sh
```

Getting an error? Read our FAQ,
[why can't I login to pico?](faq#why-cant-i-login-to-pico).

All we need to create an account is your username. This username will be used
for all of your service domains. For example, if your username is `glossy`, we
will create the following domains on your behalf:

```
glossy.prose.sh
glossy.pgs.sh
glossy.pastes.sh
```

# Next Steps

After you have created your pico account, we **highly** recommend creating an
[API token](/api-token) and keeping it someplace safe. This is how users can
recover their account if they lose their SSH private key.

What service did you want to use first?

<div class="group-h justify-center items-center my-4">
  <div class="box flex justify-center items-center" style="height: 100px; width: 150px;">
    <a href="/pgs" class="flex flex-col items-center">
      <img width="40" src="/logo-pgs.svg" />
      pages
    </a>
  </div>

<div class="box flex justify-center items-center" style="height: 100px; width: 150px;">
    <a href="/tuns" class="flex flex-col items-center">
      <img width="40" src="/logo-tuns.svg" />
      tuns
    </a>
  </div>

<div class="box flex justify-center items-center" style="height: 100px; width: 150px;">
    <a href="/pipe" class="flex flex-col items-center">
      <img width="40" src="/logo-pipe.svg" />
      pipe
    </a>
  </div>

<div class="box flex justify-center items-center" style="height: 100px; width: 150px;">
    <a href="/prose" class="flex flex-col items-center">
      <img width="40" src="/logo-prose.svg" />
      prose
    </a>
  </div>

<div class="box flex justify-center items-center" style="height: 100px; width: 150px;">
    <a href="/feeds" class="flex flex-col items-center">
      rss-to-email
    </a>
  </div>
</div>
