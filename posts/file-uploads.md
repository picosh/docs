---
title: File uploads
description: How to upload and download files from pico services
keywords: [pico, file, upload, download]
toc: true
---

All of our services require users to send us files in order to manage content.
Read the [How it Works](/how-it-works) for an under-the-hood, technical summary.

# How do I upload files?

Unless otherwise specified, all our services support the following ways to
upload files.

## rsync

```bash
rsync hello-world.md {service}:/
```

For pgs:

```bash
rsync -rv public/ {service}:/site/
```

### What rsync options are supported?

Because in our Go SSH server we re-implement `rsync`, many options are currently
not supported. For example, `--dry-run` is not supported. At
this time, the only options we supported are the following:

- `-r`
- `-v`
- `--delete`

## scp

There are two versions of `scp`, depending on your openssh version. Anything
`< v9.0` used "legacy scp." Anything `>= v9.0` uses `sftp`.

```bash
scp hello-world.md {service}:/
```

for pgs:

```bash
scp -R public/ {service}:/site/
```

## sftp

```bash
sftp {service}

sftp> ls
hello-world.md

sftp> rm hello-world.md

sftp> put hello-world.md
```

```bash
echo 'put hello-world.md' | sftp {service}
```

## sshfs

Requirement: [sshfs](https://github.com/libfuse/sshfs)

`sshfs` will allow users to mount their blog and sites like any other drive. So
you'll be able to view, edit, create, remove, and move folders and files like a
normal filesystem!

Some use cases we think are impactful:

- Debug production sites
- Run cli commands on your production sites
- Grep/find files across all your sites
- Create a development site that you use as a pgs dev server
- Make quick edits to a blog post live
- Run a formatter on your blog posts
- Easier and faster than git-ops (add+commit+push+wait-for-cicd)

### blog with prose

mount your prose.sh blog:

```bash
mkdir ~/blog
sshfs prose.sh:/ ~/blog
# edit files using your favorite editor
nvim ~/blog/hello-world.md
# changes are published live!

# unmount
umount ~/blog
```

### sites with pages

mount your pgs.sh sites:

```bash
mkdir ~/sites
sshfs pgs.sh:/ ~/sites
# edit files using your favorite editor
nvim ~/sites/myproj/index.html
# changes are published live!
```

mount a single site:

```bash
# image you have a static-site builder
cd ~/my_site
# mount your ssg's output folder
sshfs pgs.sh:/my_site ./public
# edit files using your favorite editor
nvim tmpl/base.html
# run ssg build command
# changes are published live!
```

So what's the downside? Well it's a little slower than a hard drive on your
machine. We are still experimenting with the technology so quirks or bugs might
come up. We would love to get your feedback.

# How do I update files?

Just send us the files you want to update. With [pgs.sh](/pgs) you can upload
single files to projects, but we also support "deploying" static sites with
[promotion and rollback](/pgs#project-promotion-and-rollback).

# How do I delete files?

The easiest way to delete a file is via `sftp`.

```bash
sftp {service}
sftp> rm hello-world.md
```

You could also mount our services via `sshfs` and then delete it that way.

# How do I download files?

Using the same tools described [here](#how-do-i-upload-files), just reverse the
order of `src` and `dest`!

```bash
rsync prose.sh:/ .
```
