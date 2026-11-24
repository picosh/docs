---
title: Access Logs
description: See which pubkeys accessed which services
---

We provide a TUI page dedicated to viewing access logs into your pico account. This page shows every time a service recognized one of your pubkeys or certs.

An entry is written when a key is offered, before the client proves it holds the private key. Anyone with a copy of your public key or cert can cause an entry, so one doesn't mean the login succeeded. Keys that aren't on your account, expired certs and [revoked certs](/access-control#revoking-certificates) aren't logged.

```bash
ssh pico.sh
# -> access-logs
```

On that page we also provide the ability to filter the logs.

![pico-logs](/pico-access-logs.png)

You can also fetch logs directly in the terminal:

```bash
ssh pico.sh -t access-logs
```

This makes it easier to pipe into tools like `jq` for more advanced viewing and filtering.
