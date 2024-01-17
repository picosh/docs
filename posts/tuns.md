---
title: tuns.sh
description: Access localhost using https
keywords: [pico, tuns]
---

Use your pico account to setup tunnels over ssh using
[sish](https://github.com/antoniomika/sish).

In hopes of making premium services more worthwhile to users, we are testing a
new service called `tuns.sh`. `tuns.sh` provides HTTP(S)/TCP/TLS tunnels to
localhost using SSH. Once you have a pico account (as setup on our other
services) and we apply a feature flag to your user, you'll be able to login to
`tuns.sh`.

> NOTICE: This service is currently in closed beta

However, anyone can get an invite!

The only requirement is to stay in our IRC channel and be willing to provide
feedback on how we can improve the service.

Join our IRC channel
[#pico.sh @ libera.chat](https://web.libera.chat/gamja?autojoin=#pico.sh) and
ask for an invite.

# Features

- Tunnels to localhost using SSH
- Create a private connection from a remote service to your localhost

# HTTP(S) Tunnels

```bash
$ ssh -p 2222 -R 80:httpbin.org:80 tuns.sh
Press Ctrl-C to close the session.

The subdomain localhost.tuns.sh is unavailable. Assigning a random subdomain.
Starting SSH Forwarding service for http:80. Forwarded connections can be accessed via the following methods:
Service console can be accessed here: https://flb.tuns.sh/_sish/console?x-authorization=[REDACTED]
HTTP: http://flb.tuns.sh
HTTPS: https://flb.tuns.sh
```

I can then access `http(s)://flb.tuns.sh` which will forward http requests to
`httpbin.org:80`. If I'm running a local webserver (like
`python3 -m http.server 8080`), I can replace `httpbin.org:80` with
`localhost:8080` and that address will forward to the http server I've just
started.

HTTP(S) tunnels also support
[custom domains](https://github.com/antoniomika/sish#custom-domains).

# TCP Tunnels

```bash
$ ssh -p 2222 -R 10001:httpbin.org:80 tuns.sh
Press Ctrl-C to close the session.

Starting SSH Forwarding service for tcp:10001. Forwarded connections can be accessed via the following methods:
TCP: tuns.sh:10001
```

Which will allow me to access http://tuns.sh:10001 (or any other tcp service,
`httpbin.org:80` just happens to be an HTTP server)

# And so much more

That's just the beginning of what tuns.sh can do. Under the hood we're using a
project
[Antonio's been working on for a few years](https://github.com/antoniomika/sish).
There's a lot we can do here!
