---
title: Web Tunnels
description: Passwordless authentication for the web
---

<iframe width="100%" height="315" src="https://www.youtube.com/embed/mgc5Ux1Srbc?si=7sILn0hEH8W5XqSk" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>

> [What are Webtunnels?](https://youtu.be/mgc5Ux1Srbc)

By leveraging SSH tunnels, we can use pubkey cryptography to authenticate users on the web. No passwords, no JWTs, no bearer tokens, no complicated webauthn, and no passkeys. All we need is an SSH keypair.

SSH port forwarding, also known as SSH tunneling, is the process of transmitting data over an encrypted secure shell connection between a local and distant server. It allows users to access services that firewalls would otherwise restrict or prevent. In our case, it allows us to authenticate and authorize users using only their SSH keypair.

Here is the source code for the [tunkit library](https://github.com/picosh/tunkit).

The way it works is the implementer would create an SSH app (using tunkit) that supports port forwarding, for example:

```bash
ssh -L 1337:localhost:80 -N pico-ui@pgs.sh
```

Then the user can access the site using `http://localhost:1337` and it will be automatically authenticated for that user. There's no need to pass around authentication tokens because every http request through that endpoint will already be "marked" with that user's pubkey.
