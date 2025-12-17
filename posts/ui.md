---
title: UI
description: user interfaces at pico
---

# SSH TUI

With this TUI you can perform a few basic operations like: create an account, manage pubkeys, manage API tokens.

To use the TUI just SSH into our site:

```bash
ssh pico.sh
```

# SSH CLI

Some of our services have a built-in cli with our `ssh` service:

```bash
ssh pico.sh help
ssh pgs.sh help
ssh feeds.pico.sh help
ssh pipe.pico.sh help
```

# pico.sh CLI

The TUI also comes with a CLI for some commands where it makes sense to print to stdout.

```
Commands: [help, user, logs, chat]
help - this message
user - display user information (returns name, id, account created, pico+ expiration)
logs - stream user logs
chat - IRC chat (must enable pty with `-t` to the SSH command)
not-found - return all status 404 requests for a host (hostname.com [year|month])
```

For example, if you want to drop directly into our chat app without going through the tui:

```
ssh pico.sh -t chat
```

> Note: the `-t` is important since you need to provide pty

Or if you want to stream user logs so you can grep or use other pipe tools:

```
ssh pico.sh logs
```

Want to programmatically print your user information?

```
ssh pico.sh user
```

Want to print all the status 404s for your sites?

```
ssh pico.sh not-found bower.sh month # just last month or `year` for the entire year
```
