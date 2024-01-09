---
title: prose.sh
description: guide to using prose.sh 
keywords: [pico, prose]
---

## Features

- Github flavor markdown
- [Custom domains](/custom-domains)
- Looks great on any device
- Bring your own editor
- You control the source files
- Terminal workflow with no installation
- Public-key based authentication
- Use sftp to manage blog
- No ads, zero browser-based tracking
- No attempt to identify users
- No platform lock-in
- No javascript
- Subscriptions via RSS
- Minimalist design
- 100% open source

## You control the source files

Create posts using your favorite editor in plain text files.

`~/blog/hello-world.md`

```md
# hello world!

This is my first blog post.

Check out some resources:

- [pico.sh](https://pico.sh)
- [lists.sh](https://lists.sh)
- [antoniomika](https://antoniomika.me)
- [erock.io](https://erock.io)

Cya!
```

## Publish your posts with one command

When your post is ready to be published, copy the file to our server with a
familiar command:

```
rsync ~/blog/* prose.sh:/
```

We'll either create or update the posts for you.

## Terminal workflow without installation

Since we are leveraging tools you already have on your computer (`ssh` and
`rsync`), there is nothing to install.

This provides the convenience of a web app, but from inside your terminal!

## Upload images for your blog

We also support image uploading (jpg, png, gif, webp, svg). Simply upload your
images alongside your markdown files and then reference them from root `/`:

```md
---
title: hello world!
---

![profile pic](/profile.jpg)
```

```bash
rsync ~/blog/*.jpg prose.sh:/
```

## Metadata

We support adding frontmatter to the top of your markdown posts. A frontmatter
looks like the following:

```
---
title: some title!
description: this is a great description
date: 2022-06-28
tags: feature, announcement
image: og_image.jpg
card: summary # or summary_large_image
draft: true
---
```

## How can I add a footer to all of my posts?

We have a special file `_footer.md` that will be appended to every single blog
post.

There is nothing that differentiates itself from the rest of the post so it's up
to you to style it. For convenience we added an `id` to the containing element
`post-footer`.

## How can I customize my blog page?

There's a special file you can upload `_readme.md` which will allow users to add
a bio and links to their blog landing page.

```
---
title: some title!
description: this is a great description
nav:
  - google: https://google.com
  - site: https://some.site
image: og_image.jpg
card: summary # or summary_large_image
favicon: favicon.ico
aliases:
  - 2023/03/10/my-post
layout: aside # or default
---
```

## How can I change the theme of my blog?

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
```

## How can I change the layout of my blog?

Inside the `_readme.md` metadata file, there's a variable layout option that
will change the layout of your blog index page.

```
---
layout: aside # or default
---
```
