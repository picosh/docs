---
title: pastes
description: Upload code snippets using rsync, scp, and sftp
keywords: [pico, pastes]
---

> pastes.sh is free and requires an [invite](/invites) or [pico+](/plus).

The easiest pastebin on the web.

# No install

Use tools you already have installed.

```bash
rsync my-changes.patch pastes.sh:/
```

Pastes must be plain text (UTF-8) and up to 3MB each.

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

# Can I render markdown or html?

Pastes ending in `.md`, `.markdown`, `.html` or `.htm` get a "rendered" link next to "raw". The rendered page is at `https://{user}.pastes.sh/render/{filename}`:

```bash
cat report.html | ssh pastes.sh report.html
# https://{user}.pastes.sh/render/report.html
```

Rendered pages never run scripts and can't submit forms. Styles, images (including `data:` URLs) and media still load. Markdown is rendered the same way as [prose](/prose).

# Want a pastebin for any file type?

[pgs.sh:/bin](/pgs#auto-deleting-site-after-time-period) is a generic pastebin for any filetype that auto-deletes after 14 days.

<hr />
<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
