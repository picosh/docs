---
title: pastes.sh
description: A pastebin for hackers
keywords: [pico, pastes]
---

# Features

- Pastes last 90 days by default
- [Ability to set custom expiration](#how-do-i-set-expiration-date)
- [Ability to "hide" pastes](#how-do-i-unlist-a-paste)
- Bring your own editor
- Terminal workflow with no installation
- Use sftp to manage pastes
- Public-key based authentication
- No ads, zero tracking
- No javascript
- Minimalist design
- 100% open source

# Pipe Support

```bash
echo "foobar" | ssh pastes.sh

echo "foobar" | ssh pastes.sh FILENAME

# if the tty warning annoys you
echo "foobar" | ssh -T pastes.sh
```

# How do I set expiration date?

Yes. The default expiration date for a paste is 90 days. We do allow the user to
set the paste to never expire. We also allow custom duration or timestamp.

```bash
echo "foobar" | ssh pastes.sh FILENAME expires=false

echo "foobar" | ssh pastes.sh FILENAME expires=2023-12-12

echo "foobar" | ssh pastes.sh FILENAME expires=1h
```

# How do I unlist a paste?

Yes. Unlisted in this context means it does not show up on your user landing
page where we show all of your pastes. In this case, yes, you can "hide" it
using a pipe command.

```bash
echo "foobar" | ssh pastes.sh FILENAME hidden=true
```
