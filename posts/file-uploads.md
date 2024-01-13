---
title: File uploads
description: How to upload and download files from pico services
keywords: [pico, file, upload, download]
---

All of our services require users to send us files in order to manage content.
Read the [How it Works](/how-it-works) for an under-the-hood, technical summary.

## How do I upload files?

Unless otherwise specified, all our services support the following ways to
upload files.

### rsync

```bash
rsync hello-world.md {user}@{service}:/
```

### scp

```bash
scp hello-world.md {user}@{service}:/
```

### sftp

```bash
sftp {user}@{service}

sftp> ls
hello-world.md

sftp> rm hello-world.md

sftp> put hello-world.md
```

## How do I update files?

Just send us the files you want to update. With [pgs.sh](/pgs) you can upload
single files to projects, but we also support "deploying" static sites with
[promotion and rollback](/pgs#project-promotion-and-rollback).

## How do I delete files?

We have a couple ways to delete files depending on your use-case.

### sftp

The easiest way to delete a file is via [sftp](#sftp).

### 0-byte file

Because `scp` does not natively support deleting files, we didn't want to bake
that behavior into our SSH Apps.

However, if a user wants to delete a post they can delete the contents of the
file and then upload it to our server. If the file contains 0 bytes, we will
remove the post. For example, if you want to delete `delete.md` you could:

```bash
cp /dev/null delete.md
scp ./delete.md {user}@prose.sh:/
```

### CMS

Alternatively, most our services have a way to delete posts via the SSH App CMS.
For example, if you want to delete a post on prose:

```bash
ssh {user}@prose.sh
# Click "Manage Posts"
# Highlight the post you want to delete
# Press "X"
```

## How do I download files?

Using the same tools describe [here](#how-do-i-upload-files) just reverse the
order of `src` and `dest`!

```bash
rsync {user}@prose.sh:/ .
```
