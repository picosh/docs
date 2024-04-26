---
title: How it works
description: A brief description of how our services work
keywords: [pico, how, it, works]
---

The secret ingredient to all our services is how we let users publish content
without needing to install anything. We accomplish this with the SSH tools you
already have installed on your system.

Want to publish a blog post? Use rsync, scp, or sftp. Want to publish a website?
Use rsync, scp, or sftp. Want to share a code snippet with a colleague? Use
rsync, scp, or sftp. Hopefully you see the trend.

By using the SSH protocol and golang's implementation of SSH, we can create
golang binaries that interface with SSH in unique ways.

[Charm's wish](https://github.com/charmbracelet/wish) is the underlying library
we use to enable all our SSH apps to work seamlessly with SSH clients.

Whenever a user uploads a file to our SSH app, we don't actually store anything
on our VM from the user. Instead we transfer the file to our database or our
object store.

We support a few clients for file uploads:

- `scp`
- `sftp`
- `rsync`

All of these are implemented using `wish` and our
[own golang library](https://github.com/picosh/send).
