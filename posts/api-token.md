---
title: API Tokens
description: Learn how to generate a pico API token
---

API tokens are used for when we cannot leverage SSH pubkey authentication.

Common examples:

- Account recovery
- [Connecting to our IRC bouncer](/irc)
- Pico+ rss feed

![token-tui](token-tui.png)

## Usage

1. SSH into our [pico TUI](/ui#ssh-tui)
2. Select "tokens" submenu
3. Type <kbd>c</kbd> to generate a new token
4. Give it a descriptive name
5. Copy the generated token string to where you want to use it
