---
title: How it Works
description: A brief description of how our services work
keywords: [pico, how, it, works]
---

The special sauce of all of our pico services is how we let users publish
changes to their sites without needing to install anything. We accomplish this
with what is colloquelly terms "SSH Apps."

By using the SSH protocol and golang's implementation of SSH, we can create
golang binaries that interface with SSH in unique ways.

[Charm's wish](https://github.com/charmbracelet/wish) golang library the
underlying library we use to enable all our SSH apps to work seamlessly with SSH
clients.

Whenever a user uploads a file to our SSH app, we don't actually store anything
on our VM from the user. Instead we hold onto that file upload and put it inside
our database -- or object store depending on the service.

We support a few clients for file uploads:

- `scp`
- `sftp`
- `rsync`

All of these are implemented using `wish` and our
[own golang library](https://github.com/picosh/send).
