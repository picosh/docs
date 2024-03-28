---
title: prose.sh
description: A blog platform for hackers
keywords: [pico, prose]
---

The easiest way to publish blog articles on the web.

> prose.sh is a free service

# Features

- Github flavor markdown
- [Custom domains](/custom-domains#prosesh)
- Publish posts using [rsync, sftp, or scp](/file-uploads)
- Looks great on any device
- Bring your own editor
- You control the source files
- Terminal workflow with no installation
- Public-key based authentication
- No ads, zero browser-based tracking
- No attempt to identify users
- No platform lock-in
- No javascript
- Subscriptions via RSS
- Minimalist design

# You control the source files

Create posts using your favorite editor in plain text files.

`~/blog/hello-world.md`

```md
# hello world!

This is my first blog post.

Check out some resources:

- [pico.sh](https://pico.sh)
- [antoniomika](https://antoniomika.me)
- [bower.sh](https://bower.sh)

Cya!
```

# Publish your posts with one command

When your post is ready to be published, copy the file to our server with a
familiar command:

```bash
rsync ~/blog/* prose.sh:/
# - or -
scp ~/blog/* prose.sh:/
# - or -
sftp prose.sh
```

[Read more about uploading files](/file-uploads).

We'll either create or update the posts for you.

# Terminal workflow without installation

Since we are leveraging tools you already have on your computer (`ssh` and
`rsync`), there is nothing to install.

This provides the convenience of a web app, but from inside your terminal!

# Upload images for your blog

We also support [image uploading](/images). Simply upload your images alongside
your markdown files and then reference them from root `/`:

```md
---
title: hello world!
---

![profile pic](/profile.jpg)
```

```bash
rsync ~/blog/*.jpg prose.sh:/
```

When you upload an image to prose, we make it web optimized (e.g. strip exif,
convert to webp, and reduce filesize). We also support an
[image manipulation API](/images#image-manipulation)!

# How are blogs structured?

Think of your blog as a flat list of files, similar to how your blog is rendered
as a flat list of posts.

> All filenames uploaded to prose must be unique.

We do **not** support storing files in different folders. On your local machine,
you can store files however you like, but once uploaded to us, we lose all
directory structure information and only keep the filename.

For example, if you have two posts in different folders but the same filename
and try to upload them to prose, the second post will overwrite the first one.

# How can I customize my blog?

There's a special file you can upload `_readme.md` which will allow users to add
a bio and links to their blog landing page.

```md
---
title: some title!
description: this is a great description
nav:
  - google: https://google.com
  - site: https://some.site
image: og_image.jpg
card: summary # or summary_large_image
favicon: favicon.ico
layout: aside # or default
---
```

## `title`

Title changes your blog's name which will show up as the metadata title for the
blog index page and changes the blog post link back to the main blog page.

## `description`

The description is used as the metadata description on the blog page.

## `nav`

This property sets the navigation links on the blog index page. The `rss` link
is always appended to the end of this list.

## `image` and `card`

These properties are used for og and image metadata when for when a blog is
shared. If a blog post does not have its own `image` or `card` properties, then
it will assume the blog's `image` and `card` properties.

## `favicon`

This property will change the `favicon` for the blog and posts. Link to an ico
image.

## `layout`

This property changes the layout of the blog index page. The options are:

- `aside`
- `default`

# How can I customize a blog post?

We support adding frontmatter to the top of your markdown posts. A frontmatter
looks like the following:

```md
---
title: some title!
description: this is a great description
date: 2022-06-28
tags: [feature, announcement]
image: og_image.jpg
card: summary # or summary_large_image
draft: true
toc: true # or false or integer for max heading depth
aliases:
  - 2023/03/10/my-post
---
```

## `title`

Title changes the title of the blog post.

## `description`

Description changes the description metadata for the blog post.

## `date`

This is the published date. It ought to be in `YYYY-MM-DD` format. If this date
is set to the future, the post will be unlisted until the date provided is today
or in the past.

## `image` and `card`

These properties are used for og and image metadata when for when a blog is
shared. If a blog post does not have its own `image` or `card` properties, then
it will assume the blog's `image` and `card` properties.

## `draft`

This property will change the listing status of a blog post. If `draft: true`
then the post will be published but unlisted. It will not show up on the blog
index page or RSS feed for your blog.

## `toc`

This property adds a table of contents to the blog post based on the headers. If
set to `false` then no table will be rendered. If set to `true` then table will
be rendered. If set to an integer greater than `0` then it will set the max
heading depth to that value. So if set to `2` the table will only render links
of heading `h1` and `h2`.

Examples:

```md
---
toc: false # no table of contents
toc: true # table of contents with all headings
toc: 1 # with only h1 headings
toc: 2 # with only h1 and h2 headings
toc: 3 # with only h1, h2, and h3 headings
---
```

## `aliases`

This property will redirect all routes listed to this blog post. Primarily used
for migrating from a different blog platform to prose.

```md
---
aliases:
  - 2023/03/10/my-post
  - my-post.html
---
```

## Without frontmatter

If the user does not provide a frontmatter, we will set the following defaults:

- `date` set to date when uploaded
- `title` inferred from first h1 header (e.g. `# header`) or else the filename
- `tags` inferred from any hashtags inside content

# Unlisted posts

When a post has a date in the future or `draft: true` then the post will be
**unlisted**. This means that it will **not** be shown on your blog page nor
will it show up in your built-in RSS feed.

However, the post is still publicly accessible! This gives users the ability to
share posts before "publishing" them.

# How can I add a footer to all of my posts?

We have a special file `_footer.md` that will be appended to every single blog
post.

There is nothing that differentiates itself from the rest of the post so it's up
to you to style it. For convenience we added an `id` to the containing element
`post-footer`.

# How can I change the theme of my blog?

There's a special file you can upload `_styles.css` which will allow users to
add a CSS file to their page. It will be the final CSS file loaded on the page
so it will overwrite whatever styles have previously been added. We've also
added a couple of convenience id's attached to the body element for the blog and
post pages.

```css
/* _styles.css */
#post {
    color: green;
}

#blog {
    color: tomato;
}

header {}
main {}
footer {}
article {}
```

Further we add the post slug as a class on the `body` element. So if you upload
a file called `barrel.md` then the body element will have:

```html
<body id="post" class="barrel"></body>
```

# How can I change the default 404 page?

Upload a `_404.md` which is formatted just like all the other posts.

```md
---
title: Not found
description: Where are you going?
---

This page doesn't exist.
```

# Ready to join pico?

<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
