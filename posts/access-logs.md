---
title: Access Logs
description: See which pubkeys accessed which services
---

We provide a TUI page dedicated to viewing access logs into your pico account. This page will show you everytime a pubkey successfully authenticated with a service in your pico account.

```bash
ssh pico.sh
# -> access_logs
```

On that page we also provide the ability to filter the logs.

![pico-logs](/pico-access-logs.png)

You can also fetch logs directly in the terminal:

```bash
ssh pico.sh -t access_logs
```

This makes it easier to pipe into tools like `jq` for more advanced viewing and filtering.
