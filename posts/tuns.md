---
title: tuns.sh
description: Access localhost using https
keywords: [pico, tuns]
---

An ngrok alternative using just SSH.

> NOTICE: This is a premium [pico+](/plus) service

# Features

- Tunnels to localhost using SSH
- Share your local webserver privately with another user

# Usage

[Read the sish docs](https://docs.ssi.sh)

```bash
# if you have a local webserver on localhost:8000:
ssh -R dev:80:localhost:8000 tuns.sh
# now anyone can access it at https://dev.tuns.sh
```

# Get an Invite

Anyone can get an invite!

The only requirement is to stay in our IRC channel and be willing to provide
feedback on how we can improve the service.

Join our [IRC channel](/irc) and ask for an invite.

# Ready to join pico?

<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
