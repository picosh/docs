---
title: rss-to-email
description: An rss email notification service
keywords: [pico, feeds]
toc: true
---

Stay up-to-date with all the RSS feeds you love.

> This is a free service

# Features

- Receive email digests for your RSS feeds
- We try to render all content within the feed as HTML (with ability to disable
  it)
- Create 1-to-many email digests
- Set digest interval from `10min` to `30day`

# Subscribe to feeds

Use our [plain text lists spec](/plain-text-lists) to create a txt file named
`daily.txt` (as an example).

Then add your email, the digest interval, and the rss feeds for which you want
to receive email notifications.

```
=: email rss@myemail.com
=: digest_interval 1day
=> https://blog.pico.sh/rss
=> https://bower.sh/rss
```

Then copy the file to our server

```bash
rsync daily.txt feeds.pico.sh:/
```

# Privacy

We don't do anything with your email besides send an email digest. If you delete
the post containing your email address, we no longer have you email address.

However, you should read our [privacy policy](/privacy) and
[terms of service](/ops).

Posts are also not accessible by the public and we provide no endpoints to view
these posts.

# Digest interval options

```
10min
1hour
6hour
12hour
1day
7day
30day
```

# Inline content

By default we attempt to render all content within a feed as HTML inside an
email digest. Sometimes users just want us to send them the links so they can
click on it and read the content on the original website.

If you don't want to see all the content, simply add a variable to your post:

```
=: inline_content false
```

# Can I create multiple email digests?

You are free to upload as many email digests as you like, referencing different
rss feeds, emails, and digest intervals.

# Can I fetch Reddit RSS feeds?

Yes, but you need to use the RSS links from `https://old.reddit.com`. The new
reddit doesn't appear to work properly.
