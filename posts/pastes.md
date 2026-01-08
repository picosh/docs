---
title: pastes
description: Upload code snippets using rsync, scp, and sftp
keywords: [pico, pastes]
---

> pastes.sh is a free service

The easiest pastebin on the web.

# No install

Use tools you already have installed.

```bash
rsync my-changes.patch pastes.sh:/
```

# Pipe Support

```bash
echo "foobar" | ssh pastes.sh

echo "foobar" | ssh pastes.sh FILENAME

# if the tty warning annoys you
echo "foobar" | ssh -T pastes.sh
```

# How do I set expiration date?

The default expiration date for a paste is 90 days. We do allow the user to set the paste to never expire. We also allow custom duration or timestamp.

```bash
echo "foobar" | ssh pastes.sh FILENAME expires=false

echo "foobar" | ssh pastes.sh FILENAME expires=2023-12-12

echo "foobar" | ssh pastes.sh FILENAME expires=1h
```

# How do I unlist a paste?

Unlisted in this context means it does not show up on your user landing page where we show all of your pastes. In this case, yes, you can "hide" it using a pipe command.

```bash
echo "foobar" | ssh pastes.sh FILENAME hidden=true
```

<hr />
<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
