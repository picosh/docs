---
title: Web Tunnels
description: Passwordless authentication for the web
---

You ready to experience a new way to authenticate websites?

We could describe how it works, but it would be easier to show you. Run this
command:

```bash
ssh -L 5000:localhost:80 -N pico-ui@pgs.sh
```

Don't worry if the above CLI command doesn't respond with anything, that means
it worked.

Then open your browser to:
[http://localhost:5000/tunnels](http://localhost:5000/tunnels)
