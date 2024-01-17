---
title: IRC
description: How we interact with our users
keywords: [pico, irc, bouncer]
---

All our realtime communication happens through IRC at #pico.sh @ libera.chat.

We are delighted to provide a hosted IRC bouncer and web client for us on the
team as well as pico users.

> NOTE: For libera.chat you must
> [ensure you have a libera account.](https://libera.chat/guides/registration)

# Generate a login token for bouncer

![pico-token-menu](https://hey.imgs.sh/pico-token-menu.png)

- SSH into a pico service CMS (e.g. `ssh prose.sh`)
- Select "Manage Tokens" submenu
- Type "n" to generate a new token
- Save token someplace safe

# Supported Clients

Next you need to pick a client to connect to the bouncer and finish setup:

- [web](#web) - at [chat.pico.sh](https://chat.pico.sh)
- [terminal](#senpai) - senpai

# Web

We provide pico users with a self-hosted version of gamja.

## Log into [chat.pico.sh](https://chat.pico.sh)

- You'll be redirected to [auth.pico.sh](https://auth.pico.sh) which implements
  a fake oauth2 service
- Enter the token from above
- Click submit

## Back to [chat.pico.sh](https://chat.pico.sh)

![irc-remember-me](https://hey.imgs.sh/irc-remember-me/x500)

- Click "remember me" (this is important)
- You'll see an error "Cannot interact with channels and users on the bouncer
  connection. Did you mean to use a specific network?" that's okay
- Message `BouncerServ` (`/msg BouncerServ help`) to configure the bouncer

Next step is to [connect to libera.chat](#connect-to-libera) section.

# Senpai

Senpai is a modern terminal client coupled pretty tightly to `soju` development
so it's a great fit for us at pico -- we use it.

[senpai (terminal client)](https://git.sr.ht/~delthas/senpai)

## Configure senpai

Create a config file

```bash
mkdir -p ~/.config/senpai
touch ~/.config/senpai/senpai.cfg
```

Edit `senpai.cfg`

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
- Network alias `libera`
- Channel `#pico.sh`

## Message `BouncerServ`

```
/msg BouncerServ help
```

## Join a network

```
network create -addr irc.libera.chat -name libera -nick <user> -enabled false
sasl set-plain -network libera <user> <sasl-pass>
network update libera -enabled true
```

## Join pico

```
/j #pico.sh
```

That's it! Join any other channels or networks using the same method.

# References

- [pico bouncer](ircs://irc.pico.sh:6697)
- [soju man page](https://soju.im/doc/soju.1.html)
- [soju](https://git.sr.ht/~emersion/soju)
- [gamja](https://git.sr.ht/~emersion/gamja)
- [senpai](https://git.sr.ht/~delthas/senpai)
