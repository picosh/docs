---
title: IRC
description: How we interact with our users
keywords: [pico, irc, bouncer]
toc: 2
---

# Realtime chat

All our realtime communication happens through IRC: #pico.sh @
[libera.chat](https://libera.chat).

We are pretty friendly and an active group so please do not hesitate to come
hang out with us.

If you don't want to use our bouncer and just want to quickly connect to our
channel, then we recommend
[the libera-hosted web client](https://web.libera.chat/gamja?autojoin=#pico.sh).

> NOTE: For libera.chat you must
> [ensure you have a libera account.](https://libera.chat/guides/registration)

# Our public bouncer

We are delighted to provide a public soju instance available to all pico users.

## Generate a login token for bouncer

![pico-token-menu](https://hey.imgs.sh/pico-token-menu.png)

- SSH into our [pico TUI](/ui#ssh-tui)
- Select "Manage Tokens" submenu
- Type "n" to generate a new token
- Save token someplace safe

## Supported Clients

Next you need to pick a client to connect to the bouncer and finish setup:

- [ssh app](#ssh-app)
- [web](#web) - at [chat.pico.sh](https://chat.pico.sh)
- [terminal](#terminal) - senpai

## SSH App

We integrated senpai into our SSH app. If you have a pico account you can
connect via:

```bash
ssh pico.sh -t chat
```

## Web

We provide pico users with a self-hosted version of gamja.

### Log into [chat.pico.sh](https://chat.pico.sh)

- You'll be redirected to [auth.pico.sh](https://auth.pico.sh)
- Enter the token from above
- Click submit

### Back to [chat.pico.sh](https://chat.pico.sh)

![irc-remember-me](https://hey.imgs.sh/irc-remember-me/x500)

- Click "remember me" (this is important)
- You'll see an error "Cannot interact with channels and users on the bouncer
  connection. Did you mean to use a specific network?" that's okay
- Message `BouncerServ` (`/msg BouncerServ help`) to configure the bouncer

Next step is to [connect to libera.chat](#connect-to-libera) section.

## Terminal

Senpai is a modern terminal client coupled pretty tightly to `soju` development
so it's a great fit for us at pico -- we use it.

[senpai (terminal client)](https://git.sr.ht/~delthas/senpai)

### Configure senpai

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

## Connect to libera

Using:

- Network `irc.libera.chat`
- Channel `#pico.sh`

### Message `BouncerServ`

```
/msg BouncerServ help
```

### Join a network

```
network create -addr irc.libera.chat -nick <user> -enabled false
sasl set-plain -network libera <user> <sasl-pass>
network update libera -enabled true
```

### Join pico

```
/j #pico.sh
```

That's it! Join any other channels or networks using the same method.

# References

- [libera.chat](https://libera.chat)
- [pico bouncer](ircs://irc.pico.sh:6697)
- [soju man page](https://soju.im/doc/soju.1.html)
- [soju](https://git.sr.ht/~emersion/soju)
- [gamja](https://git.sr.ht/~emersion/gamja)
- [senpai](https://git.sr.ht/~delthas/senpai)
