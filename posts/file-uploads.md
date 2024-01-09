---
title: File uploads
description: How to upload and download files from pico services
keywords: [pico, file, upload, download]
---

## How do I upload files?

Unless otherwise specified, all of our services support the following ways to
upload files to our services.

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

> You do **not** need to re-upload **all** of your files everytime you make a
> change to a single post or file.

All you have to do to update a post or file is re-upload the file using the same
tools previously explained.

## How do I delete files?

### sftp

The easiest way to delete a file is via [sftp](#sftp).

### 0-byte file

Because `scp` does not natively support deleting files, we didn't want to bake
that behavior into my ssh server.

However, if a user wants to delete a post they can delete the contents of the
file and then upload it to our server. If the file contains 0 bytes, we will
remove the post. For example, if you want to delete `delete.md` you could:

```bash
cp /dev/null delete.md
scp ./delete.md prose.sh:/
```

### CMS

Alternatively, most of our services have a way to delete posts via the SSH App
CMS. For example, if you want to delete a post on prose:

```bash
ssh prose.sh
# Click "Manage Posts"
# Highlight the post you want to delete
# Press "X"
```

## How do I download files?

Using the same tools describe [here](#how-do-i-upload-files) just reverse the
order of `src` and `dest`!
