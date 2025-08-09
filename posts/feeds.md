---
title: rss-to-email
description: Receive email digests for your RSS feeds using SSH
keywords: [pico, feeds]
toc: 1
---

Stay up-to-date with all the RSS feeds you love.

> This is a free service

# Features

- Receive email digests for your RSS feeds
- We try to render all content within the feed as HTML (with ability to disable
  it)
- Create 1-to-many email digests
- Set digest interval using cron format
- Only receive RSS feed items that haven't already been sent

# Subscribe to feeds

Use our [plain text lists spec](/plain-text-lists) to create a txt file named
`daily.txt` (as an example).

Then add your email, the digest interval, and the rss feeds for which you want
to receive email notifications.

```
=: email rss@myemail.com
=: cron 0 13 * * *
=> https://blog.pico.sh/rss
=> https://bower.sh/rss
```

Then copy the file to our server

```bash
rsync daily.txt feeds.pico.sh:/
# or
scp daily.txt feeds.pico.sh:/
```

# cron

Control when a digest is sent using the popular cron format.

> We don't support cron's less than every 5 minutes: `*/5 * * * *`

See docs for the cron impl we use:
[adhocore/gronx](https://github.com/adhocore/gronx?tab=readme-ov-file#cron-expression)

Some examples:

```
cron */10 * * * * # every 10 minutes
cron 0 * * * *    # every hour
cron 0 13 * * *   # every day at 13:00 utc
cron 0 13 1 * *   # every first day of the month at 13:00 utc
```

# Inline content

By default we attempt to render all content within a feed as HTML inside an
email digest. Sometimes users just want us to send them the links so they can
click on it and read the content on the original website.

If you don't want to see all the content, simply add a variable to your post:

```
=: inline_content false
```

# Remote CLI

We provide a remote CLI that users can use to perform a couple of simple but
powerful operations

```bash
$ ssh feeds.pico.sh help

Commands: [help, ls, rm, run]

Cmd            Desc
help           this help text
ls             list feed digest posts with metadata
rm {filename}  removes feed digest post
run {filename} runs the feed digest post immediately, ignoring cron timer
```

# Debugging

Sometimes a feed digest post will fail to send an email digest. Here are some
common reasons:

- Not time to fetch RSS yet
- All RSS feed entries have already been fetched
- Feed digest post validation error
- Requests fail because of network connectivity issues
- RSS feed is not accessible (e.g. site admins change the url or discontinue it)
- We are perceived as a bot and are blocked

We have two primary ways for a customer to debug why their email digests are not
being sent: [logs](/logs) and the feeds [remote cli](#remote-cli).

We provide the ability for users to immediately run their feed post, ignoring
digest interval validation. This will also show all logs related to fetching
that post.

```bash
ssh feeds.pico.sh run {filename}
```

If you run this command and still do not understand why your email digest isn't
being sent, then please [contact us](/contact).

# Keep Alive

We require the user to click a link in their email digest every once in awhile
in order to keep the feed digest post active. This user requirement plays an
important role in keeping our rss-to-email service healthy. If we didn't require
user interaction then we could keep sending an email digest to users that never
read their digests.

This keep-alive interval is subject to change. As of writing this (02/2025) we
require the user to click the keep-alive link once every **12 months**.

# Multiple feed digest posts

You are free to upload as many email digests as you like, referencing different
rss feeds, emails, and digest intervals.

# Remove feed digest post

There are two main options for removing a post: [remote CLI](#remote-cli) and
[sftp](/file-uploads#how-do-i-delete-files).

# Reddit

Yes, but you need to use the RSS links from `https://old.reddit.com`. The new
reddit doesn't appear to work properly.

# Privacy

We don't do anything with your email besides send an email digest. If you delete
the post containing your email address, we no longer have you email address.

However, you should read our [privacy policy](/privacy) and
[terms of service](/ops).

Posts are also not accessible by the public and we provide no endpoints to view
these posts.
