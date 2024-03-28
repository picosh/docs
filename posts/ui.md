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

This is an experimental web-based UI leveraging [web tunnels](/tunnels).

To use the web UI, create an SSH local forward connection to our [pgs](/pgs)
site:

```bash
ssh -L 1337:localhost:80 -N pico-ui@pgs.sh
```

Then open your browser and navigate to [localhost:1337](http://localhost:1337).

## SSH Config

The SSH tunnel command can be quite a lot to remember if you aren't using it
consistently everyday. Instead, you can setup an SSH config entry to do all the
work for you, here is an example config entry inside `~/.ssh/config`:

```
Host ui
    User pico-ui
    Hostname pgs.sh
    LocalForward 0.0.0.0:1337 localhost:80
    SessionType none
```

## imgs

If you want to view your docker repositories on [imgs.sh](https://pico/imgs.sh)
then you need to open an additional SSH tunnel:

```bash
ssh -L 1338:localhost:80 -N imgs.sh
```
