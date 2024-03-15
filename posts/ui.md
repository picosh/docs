---
title: UI
description: user interfaces at pico
---

We have two user interfaces: SSH TUI and a Web UI.

# SSH TUI

With this TUI you can perform a few basic operations like: create an account,
manage pubkeys, manage API tokens.

To use the TUI just SSH into our site:

```bash
ssh pico.sh
```

# Web UI

This is an experimental web-based UI leverage [web tunnels](/tunnels).

To use the web UI, create an SSH local forward connection to our [pgs](/pgs)
site:

```bash
ssh -L 1337:localhost:80 -N pico-ui@pgs.sh
```

Then open your browser and navigate to [localhost:1337](http://localhost:1337).
