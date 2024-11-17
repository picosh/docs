---
title: tuns
description: Host public web services on localhost
keywords: [pico, tuns]
---

An `ngrok` alternative using just SSH.

> NOTICE: This is a premium [pico+](/plus) service

# Features

- A zero-install developer tool
- Host public web services on `localhost`
- Host public tcp services on `localhost`
- Share your local webserver privately with another user
- Managed [sish](https://docs.ssi.sh) service

Using SSH tunnels, we can forward requests to your localhost from https, wss,
and tcp.

![sish-example](https://docs.ssi.sh/hiw-sish-public.png)

## Demo

<iframe width="100%" height="315" src="https://www.youtube.com/embed/wZHiuR9RqGw?si=AJLBbg5jc8ET0lvB" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>

> [tuns demo](https://www.youtube.com/watch?v=wZHiuR9RqGw)

# Use cases

Think of tuns as a developer tool. It is designed for the individual developer
who is looking to prototype, demo, or otherwise deploy services without the
overhead of managing a production environment with TLS, HTTP reverse proxy, and
provisioning cloud infrastructure.

By using tuns you get automatic and public https for local web services.

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

# User namespace

When creating a tunnel to `tuns` we always prefix the name with your username:

```
{user}-{name}.tuns.sh
```

This includes when a client is using tuns as a ProxyJump:

```bash
ssh -R foobar:22:localhost:22 tuns.sh
# On the client side
ssh -J tuns.sh {user}-foobar
```

# tunmgr

A tunnel manager for docker services.

tunmgr automatically set's up tunnels for docker services. It utilizes `Expose`
ports as well as DNSNames (and the container name/id) to setup different
permutations of tunnels.

> [source code](https://github.com/picosh/tunmgr)

```yml
services:
  tunmgr:
    image: ghcr.io/picosh/tunmgr:latest
    restart: always
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - $HOME/.ssh/id_ed25519:/key:ro
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/health"]
      interval: 2s
      timeout: 5s
      retries: 5
      start_period: 1s
  httpbin:
    image: kennethreitz/httpbin
    depends_on:
      tunmgr:
        condition: service_healthy
    # labels: # or provide tunnel names and ports explicitly
    #   tunmgr.names: httpbin # Comma separated list of names. Can be an empty
    #   tunmgr.ports: 80:80 $ Comma separated list of port maps. (remote:local)
    command: gunicorn -b 0.0.0.0:80 httpbin:app -k gevent --access-logfile -
```

With that docker compose file `httpbin` will be exposed as a public service on
tuns.

# How do I keep a tunnel open?

If you don't want to use `tunmgr` then we highly recommend using a tool like
[autossh](https://github.com/Autossh/autossh) to automatically restart a SSH
tunnel if it exits.

```bash
autossh -M 0 -R dev:80:localhost:8000 tuns.sh
```

# UDP Tunneling

Using `tuns`, you have the ability to tunnel UDP traffic without any external
binary, meaning all using SSH. This makes use of the SSH tunneling
functionality. To get started, you need to follow a few steps:

1. Start some UDP service that you want to forward. For example, a simple socat
   echo server:

   ```bash
   socat -v PIPE udp-recvfrom:5553,fork
   ```

2. SSH into tuns requesting a tun/tap with the information of where the service
   is running. This needs to be done as root. Replace
   `local-ip-of-machines-main-interface` with the ip address of the main
   interface for proper routing.

   ```bash
   sudo ssh -w 0:0 tuns.sh \
     udp-forward=10000:local-ip-of-machines-main-interface:5553
   ```

3. Bring the tunnel interface up and assign an ip that is link local (also as
   root):

   ```bash
   ip link set tun0 up; ip r a 10.1.0.1 dev tun0
   ```

4. Start a udp client to tuns.sh:10000. Here's one with netcat:

   ```bash
   nc -u tuns.sh 10000
   ```

<hr />
<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
