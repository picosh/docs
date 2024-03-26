---
title: Web Tunnels
description: Passwordless authentication for the web
---

Are you ready to experience a new way to authenticate users on your website?

<iframe width="560" height="315" src="https://www.youtube.com/embed/pRLC7dzxZzc?si=CR3W0v9vaRTbQAee" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>

[![Introducing Web Tunnels](http://img.youtube.com/vi/pRLC7dzxZzc/0.jpg)](https://youtu.be/pRLC7dzxZzc?si=CR3W0v9vaRTbQAee)

By leveraging SSH tunnels, we can use pubkey cryptography to authenticate users
on the web. No passwords, no JWTs, no bearer tokens, no complicated webauthn,
and no passkeys. All we need is an SSH keypair.

SSH port forwarding, also known as SSH tunneling, is the process of transmitting
data over an encrypted secure shell connection between a local and distant
server. It allows users to access services that firewalls would otherwise
restrict or prevent. In our case, it allows us to authenticate and authorize
users using only their SSH keypair.

Ready for a demo? Run this command to access our pico web UI:

```bash
ssh -L 1337:localhost:80 -N pico-ui@pgs.sh
```

What is this command? This command initiates an SSH local port forward which
redirects traffic from a local port on the client machine to a specified port on
a remote server via an SSH connection. In this case, it connects to our private
site hosted on [pgs.sh](/pgs).

Don't worry if the above CLI command doesn't respond with anything, that means
it worked.

Then open your browser:

<a href="http://localhost:1337/tunnels" class="btn-link">localhost:1337/tunnels</a>
