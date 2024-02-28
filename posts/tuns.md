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

# Docs

[Read the sish docs](https://docs.ssi.sh)

# Usage

```bash
# if you have a local webserver on localhost:8000:
ssh -R dev:80:localhost:8000 tuns.sh
# now anyone can access it at https://dev.tuns.sh
```

# Ready to join pico?

<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
