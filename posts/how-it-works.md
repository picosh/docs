---
title: How it works
description: A brief description of how our services work
keywords: [pico, how, it, works]
---

The secret ingredient to all our services is how we let users publish content
without needing to install anything. We accomplish this with the SSH tools you
already have installed on your system. All a user needs is an SSH client to
manage their content.

By using the SSH protocol and Go's implementation of SSH, we can create Go
binaries that interface with SSH in unique ways. Further, we are inside the
context of a Go binary, not a traditional SSH session where the user could
figure out how to execute arbitrary commands.

[charm.sh wish](https://github.com/charmbracelet/wish) is the underlying library
we use to enable all our SSH apps to work seamlessly with SSH clients. Wish lets
end-developers construct an SSH server with a middleware stack. This makes it
easy to programmatically compose and extend a traditional SSH server. So we have
built middleware to serve our needs:

- Authentication and authorization with keypairs and a database
- Uploading files with `scp`, `sftp`, `rsync`, and `sshfs`
- Piping into an SSH server
- Rendering a TUI
- Remote CLI

All of these are just middleware. After understanding the features and
limitations of the SSH protocol and implementing middleware with `wish`, all
that's left is a traditional Go app.

Whenever a user uploads a file to our SSH app, we don't actually store anything
in our VM from the user. Instead we transfer the file, in-memory, to a database
or object store. This makes our SSH apps stateless.

We support a few clients for file uploads:

- `scp`
- `sftp`
- `rsync`
- `sshfs`

All of these are implemented using `wish` and our
[own Go library](https://github.com/picosh/send).
