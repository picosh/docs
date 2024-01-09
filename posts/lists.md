---
title: lists.sh
description: guide to using lists.sh 
keywords: [pico, lists]
---

## Features

- Just lists
- Looks great on any device
- Bring your own editor
- You control the source files
- Terminal workflow with no installation
- Public-key based authentication
- No ads, zero browser-based tracking
- No platform lock-in
- No javascript
- Subscriptions via RSS
- Not a platform for todos
- Minimalist design
- 100% open source

## Publish your posts with one command

When your post is ready to be published, copy the file to our server with a
familiar command:

```bash
rsync ~/lists/*.txt lists.sh:/
```

We'll either create or update the lists for you.

## Plain text format

A [simple specification](/plain-text-lists) that is flexible and with no frills.

## How do I change my blog's name?

All you have to do is create a post titled `_header.txt` and add some
information to the list.

```
=: title My new blog!
=: description My blog description!
=> https://xyz.com website
=> https://twitter.com/xyz twitter
```

## How do I add an introduction to my blog?

All you have to do is create a post titled `_readme.txt` and add some
information to the list.

```
=: list_type none
# Hi my name is Bob!
I like to sing. Dance. And I like to have fun fun fun!
```

Whatever is inside the `_readme` file will get rendered (as a list) right above
your blog posts. Neat!

## How can I change the layout of my blog?

Inside the `_header.txt` metadata file, there's a variable `layout` option that
will change the layout of your blog index page.

Currently supported options:

```
=: layout default
=: layout aside
```
