---
title: Web Tunnels
description: Passwordless authentication for the web
---

Are you ready to experience a new way to authenticate users on your website?

We could describe how it works, but it would be easier to show you. Run this
command:

```bash
ssh -L 1337:localhost:80 -N pico-ui@pgs.sh
```

Don't worry if the above CLI command doesn't respond with anything, that means
it worked.

Then open your browser to:
[http://localhost:1337/tunnels](http://localhost:1337/tunnels)
