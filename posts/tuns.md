---
title: tuns.sh
description: Access localhost using https
keywords: [pico, tuns]
---

An ngrok alternative using just SSH.

> NOTICE: This is a premium [pico+](/plus) service

# Features

- Managed [sish](https://docs.ssi.sh) service
- Tunnels to localhost using SSH
- Share your local webserver privately with another user

Using SSH tunnels, we can forward requests to your localhost from https, wss,
and tcp.

![sish-example](https://docs.ssi.sh/hiw-sish-public.png)

# Docs

[Read the sish docs](https://docs.ssi.sh)

# Example Usage

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
