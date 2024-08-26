---
title: tuns
description: Access localhost using https
keywords: [pico, tuns]
---

An `ngrok` alternative using just SSH.

> NOTICE: This is a premium [pico+](/plus) service

# Features

- Managed [sish](https://docs.ssi.sh) service
- A zero-install developer tool
- Host publicly available web services on `localhost`
- Host publicly available TCP services on `localhost`
- Share your local webserver privately with another user

Using SSH tunnels, we can forward requests to your localhost from https, wss,
and tcp.

![sish-example](https://docs.ssi.sh/hiw-sish-public.png)

## Demo

<iframe width="560" height="315" src="https://www.youtube.com/embed/wZHiuR9RqGw?si=AJLBbg5jc8ET0lvB" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>

> [tuns demo](https://www.youtube.com/watch?v=wZHiuR9RqGw)

# Use cases

Think of tuns as a developer tool. It is designed for the individual developer
who is looking to prototype, demo, or otherwise deploy services without the
overhead of managing a production environment with TLS, HTTP reverse proxy, and
provisioning cloud infrastructure.

Our tuns service can be used as a developer tool wherever https is required.
Need TLS in order to run something properly in a development environment? Use
tuns to get automatic TLS.

Want to prototype a web service without fully deploying it in the cloud? You can
go from starting a local web service to sharing it with the world using a single
SSH command.

Hosting public web services from your home has never been easier with tuns.

# Docs

We manage a completely separate doc site for all things related to `sish`:

- [Read the sish docs](https://docs.ssi.sh)

# Example Usage

```bash
# if you have a local webserver on localhost:8000:
ssh -R dev:80:localhost:8000 tuns.sh
# now anyone can access it at https://{user}-dev.tuns.sh
```

Ready to join pico?

<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
