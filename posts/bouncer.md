---
title: IRC Bouncer
description: pico+ users can enjoy our private bouncer
keywords: [irc, bouncer, soju]
toc: 1
---

> Our bouncer is a [pico+](/plus) service

We are delighted to provide a private soju instance available to all [pico+](/plus) users.

> NOTE: For libera.chat you must [ensure you have a libera account.](https://libera.chat/guides/registration)

# Generate a login token for bouncer

![pico-token-menu](/pico-tokens.png)

- SSH into our [pico TUI](/ui#ssh-tui)
- Select "Manage Tokens" submenu
- Type "n" to generate a new token
- Save token someplace safe

# Supported Clients

Next you need to pick a client to connect to the bouncer and finish setup:

- [ssh app](#ssh-app)
- [web](#web) - at [chat.pico.sh](https://chat.pico.sh)
- [terminal](#terminal) - senpai

# SSH App

We integrated senpai into our SSH app. If you have a pico account you can connect via:

```bash
ssh pico.sh
# -> Chat
#
# OR go there directly:
ssh pico.sh -t chat
```

# Web

We provide pico users with a self-hosted version of gamja.

## Log into [chat.pico.sh](https://chat.pico.sh)

- You'll be redirected to [auth.pico.sh](https://auth.pico.sh)
- Enter the token from above
- Click submit

## Back to [chat.pico.sh](https://chat.pico.sh)

![irc-remember-me](https://blog.pico.sh/irc-remember-me.png/x500)

- Click "remember me" (this is important)
- You'll see an error "Cannot interact with channels and users on the bouncer connection. Did you mean to use a specific network?" that's okay
- Message `BouncerServ` (`/msg BouncerServ help`) to configure the bouncer

Next step is to [connect to libera.chat](#connect-to-libera) section.

# Terminal

Senpai is a modern terminal client coupled pretty tightly to `soju` development so it's a great fit for us at pico -- we use it.

[senpai (terminal client)](https://git.sr.ht/~delthas/senpai)

## Configure senpai

Create a config file

```bash
mkdir -p ~/.config/senpai
touch ~/.config/senpai/senpai.scfg
```

Edit `senpai.scfg`

```
address ircs://irc.pico.sh:6697
nickname <user>
password "<pico-token>"
tls true
```

Open senpai

```bash
senpai
```

Next step is to [connect to libera.chat](#connect-to-libera) section.

# Connect to libera

Using:

- Network `irc.libera.chat`
- Channel `#pico.sh`

## Message `BouncerServ`

```
/msg BouncerServ help
```

## Join a network

```
network create -addr irc.libera.chat -nick <network_nick> -enabled false
sasl set-plain -network irc.libera.chat <sasl-user> <sasl-pass>
network update irc.libera.chat -enabled true
```

## Join pico

```
/j #pico.sh
```

That's it! Join any other channels or networks using the same method.

# References

- [libera.chat](https://libera.chat)
- [pico bouncer](ircs://irc.pico.sh:6697)
- [soju man page](https://soju.im/doc/soju.1.html)
- [soju](https://codeberg.org/emersion/soju)
- [gamja](https://codeberg.org/emersion/gamja)
- [senpai](https://git.sr.ht/~delthas/senpai)
