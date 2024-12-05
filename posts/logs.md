---
title: Logs
description: Learn how to view your logs
---

In an effort for users to debug their own sites, email digests, tunnels, etc. we
want to provide a log drain that users can access.

We provide a TUI page where the user to view realtime logs.

> These logs are realtime we do **not** provide historical logs so when you
> leave the page the logs are gone.

```bash
ssh pico.sh
# -> Logs
```

On that page we also provide the ability to filter the logs.

You can also fetch logs directly in the terminal:

```bash
ssh pico.sh -t logs
```

This makes it easier to pipe into tools like `jq` for more advanced viewing and
filtering.
