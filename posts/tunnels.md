---
title: Web Tunnels
description: Passwordless authentication for the web
---

Are you ready to experience a new way to authenticate users on your website?

By leveraging SSH tunnels, we can use pubkey cryptography to authenticate users
on the web. No passwords, no JWTs, no bearer tokens, no complicated webauthn,
and no passkeys. All we need is an SSH keypair.

We could describe how it works, but it would be easier to show you. Run this
command to access our pico web UI:

```bash
ssh -L 1337:localhost:80 -N pico-ui@pgs.sh
```

Don't worry if the above CLI command doesn't respond with anything, that means
it worked.

Then open your browser:

<a href="http://localhost:1337/tunnels" class="btn-link">localhost:1337/tunnels</a>
